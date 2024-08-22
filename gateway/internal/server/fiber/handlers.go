package fiber

import (
	"fmt"
	"github.com/alserov/car_insurance/gateway/internal/service"
	"github.com/alserov/car_insurance/gateway/internal/service/models"
	"github.com/alserov/car_insurance/gateway/internal/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
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

// CreateInsurance godoc
// @Summary      CreateInsurance
// @Description  create new insurance
// @Tags         insurance
// @Accept       json
// @Produce      json
// @Param        input   body      models.Insurance  true  "insurance data"
// @Success      201  {int}  0
// @Failure      400  {string}  "invalid data"
// @Failure      500  {string}  "internal error"
// @Router       /insurance/new [post]
func (i insurance) CreateInsurance(c *fiber.Ctx) error {
	var ins models.Insurance
	if err := c.BodyParser(&ins); err != nil {
		return utils.NewError(err.Error(), utils.BadRequest)
	}

	err := i.service.CreateInsurance(c.Context(), ins)
	if err != nil {
		return fmt.Errorf("failed to create insurance: %w", err)
	}

	c.Status(http.StatusCreated)

	return nil
}

// GetInsuranceData godoc
// @Summary      GetInsuranceData
// @Description  get insurance data
// @Tags         insurance
// @Accept       json
// @Produce      json
// @Param        addr   query      string  true  "account addr"
// @Success      200  {object}  models.InsuranceData
// @Failure      400  {string}  "invalid data"
// @Failure      500  {string}  "internal error"
// @Router       /insurance/info [get]
func (i insurance) GetInsuranceData(c *fiber.Ctx) error {
	addr := c.Query("addr")

	data, err := i.service.GetInsuranceData(c.Context(), addr)
	if err != nil {
		return fmt.Errorf("failed to get insurance data: %w", err)
	}

	_ = c.JSON(data)

	return nil
}

// Payoff godoc
// @Summary      Payoff
// @Description  get insurance payoff
// @Tags         insurance
// @Accept       json
// @Produce      json
// @Param        input   body      models.Payoff  true  "payoff data"
// @Success      201  {int}  0
// @Failure      400  {string}  "invalid data"
// @Failure      500  {string}  "internal error"
// @Router       /insurance/payoff [post]
func (i insurance) Payoff(c *fiber.Ctx) error {
	var pay models.Payoff
	if err := c.BodyParser(&pay); err != nil {
		return utils.NewError(err.Error(), utils.BadRequest)
	}

	err := i.service.Payoff(c.Context(), pay)
	if err != nil {
		return fmt.Errorf("failed to payoff: %w", err)
	}

	c.Status(http.StatusCreated)

	return nil
}
