package vehicle_router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	uid, err := uuid.Parse(reqUid.Uuid)
	if err != nil {
		v.logger.Errorf("Failed Parser uuid %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(util.ErrorBadRequestUuid.Error()))

	}

	err = v.service.Update(uid, req.Brand, req.Model, req.LicensePlate, req.Color, req.YearOfManufacture)
	if err != nil {
		v.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse("Failed Updated"))
		return
	}
	v.logger.Infof("Updated Successful uuid: %s", reqUid)

	ctx.JSON(http.StatusOK, req)

}
