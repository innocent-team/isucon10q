package main

import (
	"github.com/labstack/echo"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// ref: https://github.com/newrelic/go-agent/blob/d1121710a4067a482f4baea39e17052bc89d03a7/v3/newrelic/instrumentation.go#L31
func AddNewRelicContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if nrApp == nil {
			next(c)
			return nil
		}
		txn := newrelic.FromContext(c.Request().Context())
		txn.SetWebResponse(c.Response().Writer)
		txn.SetWebRequestHTTP(c.Request())

		r := newrelic.RequestWithTransactionContext(c.Request(), txn)
		c.SetRequest(r)

		next(c)
		return nil
	}
}
