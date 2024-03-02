package vehicle_router

import (
	"github.com/gin-gonic/gin"
	dto "github.com/moura95/go-ddd/internal/dtos/vehicle"
	"github.com/moura95/go-ddd/internal/infra/util"

	"net/http"
)

type vehicleReq struct {
	Brand             string `json:"brand" binding:"required"`
	Model             string `json:"model" binding:"required"`
	YearOfManufacture uint   `json:"year_of_manufacture"`
	LicensePlate      string `json:"license_plate"`
	Color             string `json:"color"`
}

func (v *VehicleRouter) create(ctx *gin.Context) {
	var req vehicleReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		v.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(400, util.ErrorResponse(util.ErrorBadRequest.Error()))
		return
	}

	ve := dto.CreateInput{
		Brand:             req.Brand,
		Model:             req.Model,
		YearOfManufacture: req.YearOfManufacture,
		LicensePlate:      req.LicensePlate,
		Color:             req.Color,
	}
	err = v.service.Create(ve)
	if err != nil {
		v.logger.Errorf("Failed Created %s", err.Error())
		ctx.JSON(500, util.ErrorResponse(util.ErrorDatabaseCreate.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, util.SuccessResponse(req))
}
