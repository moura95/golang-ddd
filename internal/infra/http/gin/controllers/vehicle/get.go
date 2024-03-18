package vehicle_router

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

type vehicleResponse struct {
	Uuid              uuid.UUID `json:"uuid"`
	Brand             string    `json:"brand"`
	Model             string    `json:"model"`
	YearOfManufacture uint      `json:"year_of_manufacture"`
	LicensePlate      string    `json:"license_plate"`
	Color             string    `json:"color"`
	DeletedAt         string    `json:"deleted_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (v *VehicleRouter) list(ctx *gin.Context) {
	vehicles, err := v.service.List()

	if err != nil {
		v.logger.Errorf("Failed Unmarshal %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(util.ErrorDatabaseRead.Error()))
		return
	}

	var resp []vehicleResponse

	for _, vehicle := range vehicles {
		resp = append(resp, vehicleResponse{
			Uuid:              vehicle.Uuid,
			Brand:             vehicle.Brand,
			Model:             vehicle.Model,
			YearOfManufacture: vehicle.YearOfManufacture,
			LicensePlate:      vehicle.LicensePlate,
			Color:             vehicle.Color,
			DeletedAt:         vehicle.DeletedAt.String,
			CreatedAt:         vehicle.CreatedAt,
			UpdatedAt:         vehicle.UpdatedAt,
		})
	}

	ctx.JSON(200, util.SuccessResponse(resp))
	return
}

func (v *VehicleRouter) getId(ctx *gin.Context) {
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
		return
	}
	vehicle, err := v.service.GetByID(uuidStr)
	if err != nil {
		v.logger.Errorf("Failed to get vehicle %s", err.Error())
		ctx.JSON(http.StatusOK, util.ErrorResponse("Not Found"))
		return
	}

	resp := vehicleResponse{
		Uuid:              vehicle.Uuid,
		Brand:             vehicle.Brand,
		Model:             vehicle.Model,
		YearOfManufacture: vehicle.YearOfManufacture,
		LicensePlate:      vehicle.LicensePlate,
		Color:             vehicle.Color,
		CreatedAt:         vehicle.CreatedAt,
		DeletedAt:         vehicle.DeletedAt.String,
		UpdatedAt:         vehicle.UpdatedAt,
	}
	v.logger.Infof("Get Id Successful uuid: %s", resp.Uuid.String())

	ctx.JSON(http.StatusOK, util.SuccessResponse(resp))

}
