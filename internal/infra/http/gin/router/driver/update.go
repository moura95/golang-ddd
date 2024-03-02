package driver_router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/dtos/driver"
	"github.com/moura95/go-ddd/internal/infra/util"
)

func (d *Driver) update(ctx *gin.Context) {
	var req driverReq
	var reqUid getIdReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(400, util.ErrorResponse(util.ErrorBadRequest.Error()))
		return
	}

	err = ctx.ShouldBindUri(&reqUid)
	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(util.ErrorBadRequestUuid.Error()))
		return
	}
	uuid, err := uuid.Parse(reqUid.Uuid)
	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(util.ErrorBadRequestUuid.Error()))
		return
	}

	updateDriver := driver.UpdateInput{
		Uuid:          uuid,
		Name:          req.Name,
		Email:         req.Email,
		TaxID:         req.TaxID,
		DriverLicense: req.DriverLicense,
		DateOfBirth:   sql.NullString{String: req.DateOfBirth},
	}

	err = d.service.Update(updateDriver)
	if err != nil {
		d.logger.Errorf("Failed Update %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, req)

}
