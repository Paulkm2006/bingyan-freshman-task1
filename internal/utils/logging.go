package utils

import (
	"bingyan-freshman-task0/internal/config"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(e *echo.Echo) {
	var zapConfig zap.Config
	if config.Config.Logger.Debug {
		zapConfig = zap.NewDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
	}
	writer, _ := rotatelogs.New(
		config.Config.Logger.Path+"%Y%m%d%H%M.log",
		rotatelogs.WithLinkName("latest.log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapConfig.EncoderConfig),
		zapcore.AddSync(writer),
		zap.DebugLevel,
	)
	Logger = zap.New(core)
	Logger.Info("logger init success")
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogMethod:   true,
		LogRemoteIP: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			Logger.Info("request",
				zap.String("Method", v.Method),
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
				zap.String("Remote IP", v.RemoteIP),
			)
			return nil
		},
	}))

}
