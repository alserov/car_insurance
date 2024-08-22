package fiber

import "github.com/gofiber/fiber/v2"

func setupRoutes(r *fiber.App, h *handler) {
	v1 := r.Group("/v1")

	ins := v1.Group("/insurance")
	ins.Get("/info", h.insurance.GetInsuranceData)
	ins.Post("/new", h.insurance.CreateInsurance)
	ins.Post("/payoff", h.insurance.Payoff)
}
