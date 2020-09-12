package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
)

func TestChairStock(t *testing.T) {
	mySQLConnectionData = NewMySQLConnectionEnv()

	var err error
	db, err = mySQLConnectionData.ConnectDB()
	if err != nil {
		t.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	defer db.Close()

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	req = req.WithContext(context.Background())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = InitChairStock(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", chairStock)

	chairs, err := chairStock.LowestChairs(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", chairs)
	t.Logf("%v %#v", chairStock.MaxLowest, chairStock.LowestID)

	err = chairStock.Buy(36, c)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", chairStock.Purged)
	err = chairStock.Buy(18, c)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", chairStock.Purged)

	chairs, err = chairStock.LowestChairs(c)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", chairStock.Purged)

}
