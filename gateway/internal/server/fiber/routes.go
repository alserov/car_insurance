package fiber

import "github.com/gofiber/fiber/v2"

func setupRoutes(r *fiber.App, h *handler) {
	v1 := r.Group("/v1")

	ins := v1.Group("/insurance")
	ins.Get("/")
	ins.Post("/")

	pay := v1.Group("/payments")
	pay.Post("/")
}
