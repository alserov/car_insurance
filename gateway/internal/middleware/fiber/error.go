package fiber

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/logger"
	"github.com/alserov/car_insurance/gateway/internal/middleware"
	"github.com/alserov/car_insurance/gateway/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func WithErrorHandler(handler fiber.Handler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := handler(ctx); err != nil {
			msg, st := utils.FromError(err)
			if st == utils.Internal {
				logger.ExtractLogger(ctx.Locals(middleware.CtxKey).(context.Context)).Error(msg)
				return nil
			}
		}

		return nil
	}
}
