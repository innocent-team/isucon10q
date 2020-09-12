package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func InCoords(c echo.Context, ctx context.Context, coordinates Coordinates) ([]Estate, error) {
	estates := []Estate{}
	b := coordinates.getBoundingBox()
	err := db.SelectContext(ctx, &estates, `
		SELECT * FROM estate
			WHERE latitude <= ? AND latitude >= ?
			AND longitude <= ? AND longitude >= ?
			AND ST_Contains(ST_PolygonFromText(?), POINT(latitude, longitude))
			ORDER BY popularity DESC, id ASC
			LIMIT ?
		`,
		b.BottomRightCorner.Latitude, b.TopLeftCorner.Latitude,
		b.BottomRightCorner.Longitude, b.TopLeftCorner.Longitude,
		coordinates.coordinatesToTextUnesc(),
		NazotteLimit)
	return estates, err
}

func OldInCoords(c echo.Context, ctx context.Context, coordinates Coordinates) ([]Estate, error) {
	b := coordinates.getBoundingBox()
	estatesInBoundingBox := []Estate{}
	query := `SELECT * FROM estate WHERE latitude <= ? AND latitude >= ? AND longitude <= ? AND longitude >= ? ORDER BY popularity DESC, id ASC`
	err := db.SelectContext(ctx, &estatesInBoundingBox, query, b.BottomRightCorner.Latitude, b.TopLeftCorner.Latitude, b.BottomRightCorner.Longitude, b.TopLeftCorner.Longitude)
	if err == sql.ErrNoRows {
		c.Echo().Logger.Infof("select * from estate where latitude ...", err)
		return nil, c.JSON(http.StatusOK, EstateSearchResponse{Count: 0, Estates: []Estate{}})
	} else if err != nil {
		c.Echo().Logger.Errorf("database execution error : %v", err)
		return nil, c.NoContent(http.StatusInternalServerError)
	}

	estatesInPolygon := []Estate{}
	for _, estate := range estatesInBoundingBox {
		validatedEstate := Estate{}

		point := fmt.Sprintf("'POINT(%f %f)'", estate.Latitude, estate.Longitude)
		query := fmt.Sprintf(`SELECT * FROM estate WHERE id = ? AND ST_Contains(ST_PolygonFromText(%s), ST_GeomFromText(%s))`, coordinates.coordinatesToText(), point)
		err = db.GetContext(ctx, &validatedEstate, query, estate.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			} else {
				c.Echo().Logger.Errorf("db access is failed on executing validate if estate is in polygon : %v", err)
				return nil, c.NoContent(http.StatusInternalServerError)
			}
		} else {
			estatesInPolygon = append(estatesInPolygon, validatedEstate)
		}
	}
	if len(estatesInPolygon) > NazotteLimit {
		return estatesInPolygon[:NazotteLimit], nil
	}
	return estatesInPolygon, nil
}

func (cs Coordinates) coordinatesToTextUnesc() string {
	points := make([]string, 0, len(cs.Coordinates))
	for _, c := range cs.Coordinates {
		points = append(points, fmt.Sprintf("%f %f", c.Latitude, c.Longitude))
	}
	return fmt.Sprintf("POLYGON((%s))", strings.Join(points, ","))
}

func (cs Coordinates) coordinatesToText() string {
	points := make([]string, 0, len(cs.Coordinates))
	for _, c := range cs.Coordinates {
		points = append(points, fmt.Sprintf("%f %f", c.Latitude, c.Longitude))
	}
	return fmt.Sprintf("'POLYGON((%s))'", strings.Join(points, ","))
}
