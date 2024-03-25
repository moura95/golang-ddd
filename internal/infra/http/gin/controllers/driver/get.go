package driver_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/infra/util"

	"net/http"
	"time"
)

type getIdReq struct {
	Uuid string `uri:"uuid" binding:"required"`
}

type driverResponse struct {
	Uuid          uuid.UUID         `json:"uuid"`
	Name          string            `json:"name"`
	Email         string            `json:"email"`
	TaxID         string            `json:"tax_id"`
	DriverLicense string            ` json:"driver_license"`
	DateOfBirth   string            `json:"date_of_birth"`
	Vehicles      []vehicleResponse `json:"vehicles"`
	DeletedAt     string            `json:"deleted_at"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}

type vehicleResponse struct {
	Uuid              uuid.UUID `repository:"uuid"`
	Brand             string    `repository:"brand"`
	Model             string    `repository:"model"`
	YearOfManufacture uint      `repository:"year_of_manufacture"`
	LicensePlate      string    `repository:"license_plate"`
	Color             string    `repository:"color"`
}

func (d *Driver) list(ctx *gin.Context) {
	drivers, err := d.service.List()
	if err != nil {
		d.logger.Errorf("Failed Get All %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse("Failed List Drivers"))
		return
	}

	var resp []driverResponse

	for _, driver := range drivers {
		resp = append(resp, driverResponse{
			Uuid:          driver.Uuid,
			Name:          driver.Name,
			Email:         driver.Email,
			TaxID:         driver.TaxID,
			DriverLicense: driver.DriverLicense,
			DateOfBirth:   driver.DateOfBirth.String,
			CreatedAt:     driver.CreatedAt,
			UpdatedAt:     driver.UpdatedAt,
		})
	}

	d.logger.Infof("Successfull List")

	ctx.JSON(200, util.SuccessResponse(resp))
	return
}

func (d *Driver) getId(ctx *gin.Context) {
	var req getIdReq

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
		return
	}

	err = uuid.Validate(req.Uuid)

	if err != nil {
		d.logger.Errorf("Uuid Invalid %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
	}

	uid, err := uuid.Parse(req.Uuid)
	if err != nil {
		d.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse("Bad Request!! Uuid: Invalid"))
		return
	}

	driver, err := d.service.GetByID(uid)
	if err != nil {
		d.logger.Errorf("Failed Get uuid %s", err.Error())
		ctx.JSON(http.StatusOK, util.SuccessResponse("Not Found"))
		return
	}
	var vehicles []vehicleResponse

	resp := driverResponse{
		Uuid:          driver.Uuid,
		Name:          driver.Name,
		Email:         driver.Email,
		TaxID:         driver.TaxID,
		DriverLicense: driver.DriverLicense,
		Vehicles:      vehicles,
		DateOfBirth:   driver.DateOfBirth.String,
		DeletedAt:     driver.DeletedAt.String,
		CreatedAt:     driver.CreatedAt,
		UpdatedAt:     driver.UpdatedAt,
	}
	for _, v := range driver.Vehicles {
		resp.Vehicles = append(resp.Vehicles, vehicleResponse{
			Uuid:              v.Uuid,
			Brand:             v.Brand,
			Model:             v.Model,
			YearOfManufacture: v.YearOfManufacture,
			Color:             v.Color,
		})
	}

	d.logger.Infof("Succesful getID: %s", resp.Uuid.String())

	ctx.JSON(http.StatusOK, util.SuccessResponse(resp))

}
