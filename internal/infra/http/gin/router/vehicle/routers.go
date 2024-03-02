package vehicle_router

import (
	"github.com/gin-gonic/gin"
	"github.com/moura95/go-ddd/internal/application/service/vehicle"
	"go.uber.org/zap"
)

type IVehicle interface {
	SetupVehicleRoute(routers *gin.RouterGroup)
}

type VehicleRouter struct {
	service vehicle.IVehicleService
	logger  *zap.SugaredLogger
}

func NewVehicleRouter(s vehicle.IVehicleService, log *zap.SugaredLogger) *VehicleRouter {
	return &VehicleRouter{
		service: s,
		logger:  log,
	}
}

func (v *VehicleRouter) SetupVehicleRoute(routers *gin.RouterGroup) {
	routers.GET("/vehicle", v.list)
	routers.GET("/vehicle/:uuid", v.getId)
	routers.PUT("/vehicle/:uuid", v.update)
	routers.DELETE("/vehicle/:uuid/hard", v.hardDelete)
	routers.DELETE("/vehicle/:uuid", v.delete)
	routers.PATCH("/vehicle/:uuid/recover", v.undelete)
	routers.POST("/vehicle", v.create)

}
