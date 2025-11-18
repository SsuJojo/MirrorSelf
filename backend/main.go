package main

import (
	"os"
	"strings"

	"mirrorself/backend/pb"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initLogger() *zap.Logger {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder, // 控制台彩色输出
		EncodeTime:    zapcore.ISO8601TimeEncoder,       // 标准时间格式
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}

	// 控制台输出
	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)

	// 文件输出
	file, err := os.OpenFile("meals.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	fileEncoder := zapcore.NewJSONEncoder(encoderCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.DebugLevel), // 控制台
		zapcore.NewCore(fileEncoder, zapcore.AddSync(file), zap.InfoLevel),          // 文件
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger
}

func main() {
	app := fiber.New()
	pb.StartPocketBase()
	logger := initLogger()
	defer logger.Sync()

	app.Static("/", "../frontend/dist")

	api := app.Group("/api")
	api.Post("/meal", func(c *fiber.Ctx) error {
		type Request struct {
			Meal string `json:"meal"`
		}
		type Response struct {
			Status string `json:"status"`
		}

			var req Request
			if err := c.BodyParser(&req); err != nil {
				logger.Error("JSON parse error", zap.Error(err))
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Cannot parse JSON",
				})
			}

		logger.Info("Received meal", zap.String("meal", req.Meal))

		return c.JSON(Response{Status: "recorded"})
	})

	api.Post("/msgsomany", func(c *fiber.Ctx) error {
		logger.Info("Someone click the button so many times~")
		return c.SendStatus(fiber.StatusOK)
	})

	app.All("/db/*", func(c *fiber.Ctx) error {
		target := "http://127.0.0.1:8090" + strings.TrimPrefix(c.OriginalURL(), "/db") // PocketBase 地址，原请求去除/db前缀
		return proxy.Do(c, target)
	})

	if err := app.Listen(":3001"); err != nil {
		logger.Fatal("Server failed", zap.Error(err))
	}
}
