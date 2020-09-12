package main

import (
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func MyWrapHandle(path string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		_, wrapped := newrelic.WrapHandle(nrApp, path, h)
		return wrapped
	}
}
