package driver_router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/infra/util"

	"net/http"
)

func (d *Driver) softDelete(ctx *gin.Context) {
	var req getIdReq

	err := ctx.ShouldBindUri(&req)

	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
		return
	}
	uuidStr, err := uuid.Parse(req.Uuid)
	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
	}

	// before hardDelete need remove relation
	err = d.service.SoftDelete(uuidStr)

	if err != nil {
		d.logger.Errorf("Failed Soft Delete %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, "Failed to HardDelete Driver")
		return
	}
	d.logger.Infof("Succesful Deleted uuid: %s", uuidStr)

	ctx.JSON(http.StatusOK, util.SuccessResponse("OK"))
}

func (d *Driver) UnDelete(ctx *gin.Context) {
	var req getIdReq

	err := ctx.ShouldBindUri(&req)

	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
		return
	}
	uuidStr, err := uuid.Parse(req.Uuid)
	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
	}

	// before hardDelete need remove relation
	err = d.service.UnDelete(uuidStr)

	if err != nil {
		d.logger.Errorf("Failed to Delete %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, "Failed to HardDelete Driver")
		return
	}

	d.logger.Infof("Succesful Delete uuid: %s", uuidStr)

	ctx.JSON(http.StatusOK, util.SuccessResponse("OK"))
}

func (d *Driver) hardDelete(ctx *gin.Context) {
	var req getIdReq

	err := ctx.ShouldBindUri(&req)

	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
		return
	}
	uuidStr, err := uuid.Parse(req.Uuid)
	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
	}

	// before hardDelete need remove relation
	err = d.service.HardDelete(uuidStr)

	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, "Failed to HardDelete Driver")
		return
	}
	d.logger.Infof("Delete Sucessful uuid: %s", uuidStr)

	ctx.JSON(http.StatusOK, util.SuccessResponse("OK"))

}
