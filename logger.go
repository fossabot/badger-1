package main

import (
	"github.com/urfave/negroni"
)

func newLoggerMiddleware(logEndpoint string) (logger *negroni.Logger) {
	logger = negroni.NewLogger()
	return
}
