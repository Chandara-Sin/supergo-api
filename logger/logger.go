package logger

import (
	"bytes"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const key = "logger"

func Middleware(log *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqID := c.Request().Header.Get("X-Request-ID")
			l := log.With(zap.String("req-id", reqID))
			c.Set(key, l)

			bodyBytes := []byte{}
			if c.Request().Body != nil {
				bodyBytes, _ = io.ReadAll(c.Request().Body)
			}

			c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			fmt.Printf("request body: %s\n", bodyBytes)

			return next(c)
		}
	}
}

func Unwrap(c echo.Context) *zap.Logger {
	val := c.Get(key)
	if log, ok := val.(*zap.Logger); ok {
		return log
	}
	return zap.NewExample()
}
