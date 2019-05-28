package main

import (
	"github.com/go-kit/kit/log"
	"time"
)

// ServiceMiddleware define service middleware
type ServiceMiddleware func(Service) Service

// loggingMiddleware Make a new type
// that contains Service interface and logger instance
type loggingMiddleware struct {
	Service
	logger log.Logger
}
// LoggingMiddleware make logging middleware
func LoggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next Service) Service {
		return loggingMiddleware{next, logger}
	}
}



func (mw loggingMiddleware) Add(a, b int) (ret int) {

	defer func(beign time.Time) {
		mw.logger.Log(
			"function", "Add",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(beign),
		)
	}(time.Now())

	ret = mw.Service.Add(a, b)
	return ret
}

