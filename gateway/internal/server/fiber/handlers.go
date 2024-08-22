package fiber

import (
	"fmt"
	"github.com/alserov/car_insurance/gateway/internal/service"
	"github.com/alserov/car_insurance/gateway/internal/service/models"
	"github.com/alserov/car_insurance/gateway/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func newHandler(srvc *service.Service) *handler {
	return &handler{
		insurance: insurance{
			service: srvc.Insurance,
		},
	}
}

type handler struct {
	insurance insurance
}

type insurance struct {
	service service.Insurance
}

func (i insurance) CreateInsurance(c *fiber.Ctx) error {
	var ins models.Insurance
	if err := c.BodyParser(&ins); err != nil {
		return utils.NewError(err.Error(), utils.BadRequest)
	}

	err := i.service.CreateInsurance(c.Context(), ins)
	if err != nil {
		return fmt.Errorf("failed to create insurance: %w", err)
	}

	return nil
}

func (i insurance) GetInsuranceData(c *fiber.Ctx) error {
	addr := c.Query("addr")

	data, err := i.service.GetInsuranceData(c.Context(), addr)
	if err != nil {
		return fmt.Errorf("failed to get insurance data: %w", err)
	}

	_ = c.JSON(data)

	return nil
}

func (i insurance) Payoff(c *fiber.Ctx) error {
	var pay models.Payoff
	if err := c.BodyParser(&pay); err != nil {
		return utils.NewError(err.Error(), utils.BadRequest)
	}

	err := i.service.Payoff(c.Context(), pay)
	if err != nil {
		return fmt.Errorf("failed to payoff: %w", err)
	}

	return nil
}
