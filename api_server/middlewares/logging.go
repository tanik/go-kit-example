package middlewares

import (
	services "go-kit-example/api_server/services"
	"time"

	"github.com/go-kit/log"
)

func LoggingMiddleware(logger log.Logger) services.ServiceMiddleware {
	return func(next services.StringService) services.StringService {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	services.StringService
}

func (mw logmw) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.StringService.Uppercase(s)
	return
}

func (mw logmw) Count(s string) (n int) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.StringService.Count(s)
	return
}
