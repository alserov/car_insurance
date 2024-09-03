package fiber

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/logger"
	"github.com/alserov/car_insurance/gateway/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func WithRecovery(handler fiber.Handler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				logger.ExtractLogger(ctx.Locals(middleware.CtxKey).(context.Context)).Error("panic recovery", logger.WithArg("error", err))
			}
		}()

		return handler(ctx)
	}
}
