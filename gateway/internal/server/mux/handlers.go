package mux

import (
	"encoding/json"
	"fmt"
	"github.com/alserov/car_insurance/gateway/internal/service"
	"github.com/alserov/car_insurance/gateway/internal/service/models"
	"github.com/alserov/car_insurance/gateway/internal/tracing"
	"github.com/alserov/car_insurance/gateway/internal/utils"
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
func (i insurance) CreateInsurance(w http.ResponseWriter, r *http.Request) error {
	ctx, span := tracing.ExtractTracer(r.Context()).Start(r.Context(), "received create insurance request")
	defer span.End()

	var ins models.Insurance
	if err := json.NewDecoder(r.Body).Decode(&ins); err != nil {
		return utils.NewError(err.Error(), utils.BadRequest)
	}

	err := i.service.CreateInsurance(ctx, ins)
	if err != nil {
		return fmt.Errorf("failed to create insurance: %w", err)
	}

	w.WriteHeader(http.StatusCreated)

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
func (i insurance) GetInsuranceData(w http.ResponseWriter, r *http.Request) error {
	ctx, span := tracing.ExtractTracer(r.Context()).Start(r.Context(), "received get insurance request")
	defer span.End()

	addr := r.URL.Query()["addr"][0]

	data, err := i.service.GetInsuranceData(ctx, addr)
	if err != nil {
		return fmt.Errorf("failed to get insurance data: %w", err)
	}

	if err = json.NewEncoder(w).Encode(data); err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

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
func (i insurance) Payoff(w http.ResponseWriter, r *http.Request) error {
	ctx, span := tracing.ExtractTracer(r.Context()).Start(r.Context(), "received payoff request")
	defer span.End()

	var pay models.Payoff
	if err := json.NewDecoder(r.Body).Decode(&pay); err != nil {
		return utils.NewError(err.Error(), utils.BadRequest)
	}

	err := i.service.Payoff(ctx, pay)
	if err != nil {
		return fmt.Errorf("failed to payoff: %w", err)
	}

	w.WriteHeader(http.StatusCreated)

	return nil
}
