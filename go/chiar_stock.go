package main

import (
	"database/sql"
	"net/http"
	"sync"

	"github.com/labstack/echo"
)

type ChairStock struct {
	mutex sync.RWMutex
	Stock map[int64]int64

	Purged    bool
	LowestID  map[int64]Chair
	Lowest    []Chair
	MaxLowest int64
}

var chairStock ChairStock

func InitChairStock(c echo.Context) error {
	chairStock = ChairStock{Stock: map[int64]int64{}, Purged: true}
	ctx := c.Request().Context()
	stocks := []struct {
		ID    int64 `db:"id"`
		Stock int64 `db:"stock"`
	}{}
	println(db, ctx, stocks)
	err := db.SelectContext(ctx, &stocks, `
		SELECT id, stock FROM chair
	`)
	if err != nil {
		c.Logger().Errorf("Init chair %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}
	chairStock.mutex.Lock()
	for _, stock := range stocks {
		chairStock.Stock[stock.ID] = stock.Stock
	}
	chairStock.mutex.Unlock()
	return nil
}

func (s *ChairStock) Buy(id int64, c echo.Context) error {
	s.mutex.RLock()
	stock, ok := s.Stock[id]
	s.mutex.RUnlock()
	if !ok || stock <= 0 {
		c.Echo().Logger.Infof("buyChair chair id \"%v\" not found", id)
		return c.NoContent(http.StatusNotFound)
	}

	s.mutex.Lock()
	s.Stock[id]--

	// if the stock < max lowest
	if s.Stock[id] <= s.MaxLowest {
		s.Purged = true
	}
	// if low stock goes 0, purge lowest
	if stock == 1 {
		_, ok = s.LowestID[id]
		if ok {
			s.Purged = true
		}
	}
	s.mutex.Unlock()

	// Thought DB (cache after)
	err := doBuyChair(int(id), c)
	if err != nil {
		return err
	}
	return nil
}

func (s *ChairStock) Add(id int64, stock int64, c echo.Context) error {
	s.mutex.Lock()
	s.Stock[id] = stock
	if stock < s.MaxLowest {
		s.Purged = true
	}
	s.mutex.Unlock()
	return nil
}

func (s *ChairStock) LowestChairs(c echo.Context) ([]Chair, error) {
	if !s.Purged {
		c.Echo().Logger.Infof("Cached Lowest Chairs")
		return s.Lowest, nil
	}
	low, err := doGetLowPriceChair(c)
	s.mutex.Lock()
	s.Lowest = low
	s.LowestID = map[int64]Chair{}
	s.MaxLowest = -1
	for _, chair := range low {
		if s.MaxLowest < chair.Stock {
			s.MaxLowest = chair.Stock
		}
		s.LowestID[chair.ID] = chair
	}
	s.Purged = false
	s.mutex.Unlock()
	return s.Lowest, err
}

func doBuyChair(id int, c echo.Context) error {
	ctx := c.Request().Context()
	tx, err := db.Beginx()
	if err != nil {
		c.Echo().Logger.Errorf("failed to create transaction : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}
	defer tx.Rollback()

	var chair Chair
	err = tx.QueryRowx("SELECT * FROM chair WHERE id = ? AND stock > 0 FOR UPDATE", id).StructScan(&chair)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Echo().Logger.Infof("buyChair chair id \"%v\" not found", id)
			return c.NoContent(http.StatusNotFound)
		}
		c.Echo().Logger.Errorf("DB Execution Error: on getting a chair by id : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	_, err = tx.ExecContext(ctx, "UPDATE chair SET stock = stock - 1 WHERE id = ?", id)
	if err != nil {
		c.Echo().Logger.Errorf("chair stock update failed : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	err = tx.Commit()
	if err != nil {
		c.Echo().Logger.Errorf("transaction commit error : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func doGetLowPriceChair(c echo.Context) ([]Chair, error) {
	ctx := c.Request().Context()
	var chairs []Chair
	query := `SELECT * FROM chair WHERE stock > 0 ORDER BY price ASC, id ASC LIMIT ?`
	err := db.SelectContext(ctx, &chairs, query, Limit)
	if err != nil {
		if err == sql.ErrNoRows {
			c.Logger().Error("getLowPricedChair not found")
			return nil, c.JSON(http.StatusOK, ChairListResponse{[]Chair{}})
		}
		c.Logger().Errorf("getLowPricedChair DB execution error : %v", err)
		return nil, c.NoContent(http.StatusInternalServerError)
	}
	return chairs, nil
}
