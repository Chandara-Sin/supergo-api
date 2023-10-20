package logger

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"math"
	"math/big"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const key = "logger"
const traceKey = "trace_id"

func MWLogger() (*zap.Logger, error) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	zapConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}

	return zapConfig.Build()
}

func ReqLoggerConfig(log *zap.Logger) middleware.RequestLoggerConfig {
	return middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogMethod:    true,
		LogHost:      true,
		LogRemoteIP:  true,
		LogUserAgent: true,
		LogError:     true,
		BeforeNextFunc: func(c echo.Context) {
			traceId := GetTraceId()
			l := log.With(zap.String(traceKey, traceId))
			c.Set(key, l)
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			l := Unwrap(c)
			l.Info("request",
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
				zap.String("method", v.Method),
				zap.String("host", v.Host),
				zap.String("remote_ip", v.RemoteIP),
				zap.String("user_agent", v.UserAgent),
				zap.NamedError("error", v.Error),
			)
			return nil
		},
	}
}

func GetTraceId() string {
	bi, _ := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(10)))),
	)
	return "SUP-" + fmt.Sprintf("%010d", bi)
}

func Unwrap(c echo.Context) *zap.Logger {
	val := c.Get(key)
	if log, ok := val.(*zap.Logger); ok {
		return log
	}
	return zap.NewExample()
}

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
