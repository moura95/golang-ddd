package driver_router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/dtos/driver"
	"github.com/moura95/go-ddd/internal/dtos/driver_vehicle"
	"github.com/moura95/go-ddd/internal/infra/util"
)

type driverReq struct {
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" binding:"required"`
	TaxID         string `json:"tax_id"`
	DriverLicense string ` json:"driver_license"`
	DateOfBirth   string `json:"date_of_birth"`
}

func (d *Driver) create(ctx *gin.Context) {
	var req driverReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(400, util.ErrorResponse("Failed Unmarshal"))

		return
	}

	dr := driver.CreateInput{
		Name:          req.Name,
		Email:         req.Email,
		TaxID:         req.TaxID,
		DriverLicense: req.DriverLicense,
		DateOfBirth:   sql.NullString{String: req.DateOfBirth},
	}

	err = d.service.Create(dr)

	if err != nil {
		d.logger.Errorf("Failed Create Driver %s", err.Error())
		ctx.JSON(500, util.ErrorResponse("Failed Create Driver"))
		return
	}
	d.logger.Infof("Create Driver succesful %s", dr.Name)

	ctx.JSON(http.StatusCreated, util.SuccessResponse(req))
}

type subscribeReq struct {
	DriverUUID  string `json:"driver_uuid"`
	VehicleUUID string `json:"vehicle_uuid"`
}

func (d *Driver) subscribe(ctx *gin.Context) {
	var req subscribeReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(400, util.ErrorResponse("Failed Unmarshal"))
		return
	}
	vehicleUUID, err := uuid.Parse(req.VehicleUUID)
	if err != nil {
		d.logger.Errorf("Failed Parser uuid %s", err.Error())
		ctx.JSON(400, util.ErrorResponse("Failed Unmarshal"))
		return
	}

	driverUUID, err := uuid.Parse(req.DriverUUID)
	if err != nil {
		d.logger.Errorf("Failed Parser Driver Uuid:%s ", err.Error())
		ctx.JSON(400, util.ErrorResponse("Failed Unmarshal"))
		return
	}

	subs := driver_vehicle.Input{
		DriverUUID:  vehicleUUID,
		VehicleUUID: driverUUID,
	}

	err = d.service.Subscribe(subs)

	if err != nil {
		d.logger.Errorf("Failed Create Relation Driver Vehicle %s", err.Error())
		ctx.JSON(500, util.ErrorResponse("Failed Create Relation Driver Vehicle"))
		return
	}
	d.logger.Infof("Create Relation Driver Vehicle succesful")

	ctx.JSON(http.StatusCreated, util.SuccessResponse(req))
}

func (d *Driver) unSubscribe(ctx *gin.Context) {
	var req subscribeReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(400, util.ErrorResponse("Failed Unmarshal"))
		return
	}
	vehicleUUID, err := uuid.Parse(req.VehicleUUID)
	if err != nil {
		d.logger.Errorf("Failed Parser uuid %s", err.Error())
		ctx.JSON(400, util.ErrorResponse("Failed Unmarshal"))
		return
	}

	driverUUID, err := uuid.Parse(req.DriverUUID)
	if err != nil {
		d.logger.Errorf("Failed Parser Driver Uuid:%s ", err.Error())
		ctx.JSON(400, util.ErrorResponse("Failed Unmarshal"))
		return
	}
	unSubs := driver_vehicle.Input{
		DriverUUID:  vehicleUUID,
		VehicleUUID: driverUUID,
	}

	err = d.service.UnSubscribe(unSubs)
	if err != nil {
		d.logger.Errorf("Failed Unsubscribe %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, "Failed Unsubscribe")
		return
	}
	d.logger.Infof("Delete Relation DriverUUID: %s VehicleUUID:%s", driverUUID, vehicleUUID)

	ctx.JSON(http.StatusOK, util.SuccessResponse("OK"))

}
