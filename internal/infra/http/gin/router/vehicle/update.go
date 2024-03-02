package vehicle_router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	dto "github.com/moura95/go-ddd/internal/dtos/vehicle"
	"github.com/moura95/go-ddd/internal/infra/util"

	"net/http"
)

func (v *VehicleRouter) update(ctx *gin.Context) {
	var req vehicleReq
	var reqUid getIdReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		v.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(400, util.ErrorResponse(util.ErrorBadRequest.Error()))
		return
	}

	err = ctx.ShouldBindUri(&reqUid)
	if err != nil {
		v.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(util.ErrorBadRequestUuid.Error()))
		return
	}
	uuidStr, err := uuid.Parse(reqUid.Uuid)
	if err != nil {
		v.logger.Errorf("Failed Parser uuid %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(util.ErrorBadRequestUuid.Error()))

	}

	updateVehicle := dto.UpdateInput{
		Uuid:              uuidStr,
		Brand:             req.Brand,
		Model:             req.Model,
		YearOfManufacture: req.YearOfManufacture,
		LicensePlate:      req.LicensePlate,
		Color:             req.Color,
	}

	err = v.service.Update(updateVehicle)
	if err != nil {
		v.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse("Failed Updated"))
		return
	}
	v.logger.Infof("Updated Successful uuid: %s", updateVehicle.Uuid.String())

	ctx.JSON(http.StatusOK, req)

}
