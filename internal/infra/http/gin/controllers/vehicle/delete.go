package vehicle_router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/infra/util"

	"net/http"
)

func (v *VehicleRouter) hardDelete(ctx *gin.Context) {
	var req getIdReq

	err := ctx.ShouldBindUri(&req)

	if err != nil {
		v.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
		return
	}
	uuidStr, err := uuid.Parse(req.Uuid)
	if err != nil {
		v.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
	}

	// before hardDelete need remove relation
	err = v.service.HardDelete(uuidStr)

	if err != nil {
		v.logger.Errorf("Failed Hard Delete %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, "Failed to HardDelete Driver")
		return
	}

	v.logger.Infof("Failed Unmarshal %s", err.Error())

	ctx.JSON(http.StatusOK, util.SuccessResponse("OK"))

}

func (v *VehicleRouter) delete(ctx *gin.Context) {
	var req getIdReq

	err := ctx.ShouldBindUri(&req)

	if err != nil {
		v.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
		return
	}
	uuidStr, err := uuid.Parse(req.Uuid)

	if err != nil {
		v.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
	}

	// before hardDelete need remove relation
	err = v.service.SoftDelete(uuidStr)

	if err != nil {
		v.logger.Errorf("Failed Soft Delete %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, "Failed to Delete Driver")
		return
	}
	v.logger.Infof("Delete Successful uuid: %s", uuidStr)

	ctx.JSON(http.StatusOK, util.SuccessResponse("OK"))

}

func (v *VehicleRouter) undelete(ctx *gin.Context) {
	var req getIdReq

	err := ctx.ShouldBindUri(&req)

	if err != nil {
		v.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
		return
	}
	uuidStr, err := uuid.Parse(req.Uuid)
	if err != nil {
		v.logger.Errorf("Failed Parser uuid %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
		return
	}

	// before hardDelete need remove relation
	err = v.service.UnDelete(uuidStr)

	if err != nil {
		v.logger.Errorf("Failed Delete %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, "Failed to Delete Driver")
		return
	}

	v.logger.Infof("delete successful uuid: %s", uuidStr)

	ctx.JSON(http.StatusOK, util.SuccessResponse("OK"))

}
