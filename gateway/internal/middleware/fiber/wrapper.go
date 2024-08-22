package fiber

import (
	"context"
	"github.com/alserov/car_insurance/gateway/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func WithWrappers(wrs ...middleware.Wrapper) func(handler fiber.Handler) fiber.Handler {
	return func(handler fiber.Handler) fiber.Handler {
		return func(c *fiber.Ctx) error {
			ctx := context.Background()

			for _, wr := range wrs {
				ctx = wr(ctx)
			}

			c.Locals(middleware.CtxKey, ctx)

			return handler(c)
		}
	}
}
