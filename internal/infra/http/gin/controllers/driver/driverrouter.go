package driver_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/moura95/go-ddd/internal/application/services/driver"
	"go.uber.org/zap"
)

type IDriver interface {
	SetupDriverRoute(routers *gin.RouterGroup)
}

type Driver struct {
	service *driverservice.Service
	logger  *zap.SugaredLogger
}

func NewDriverRouter(s *driverservice.Service, log *zap.SugaredLogger) *Driver {
	return &Driver{
		service: s,
		logger:  log,
	}
}

func (d *Driver) SetupDriverRoute(routers *gin.RouterGroup) {
	routers.GET("/driver", d.list)
	routers.GET("/driver/:uuid", d.getId)
	routers.PUT("/driver/:uuid", d.update)
	routers.DELETE("/driver/:uuid/hard", d.hardDelete)
	routers.DELETE("/driver/:uuid", d.softDelete)
	routers.PATCH("/driver/:uuid/recover", d.UnDelete)
	routers.POST("/driver", d.create)
	routers.POST("/driver/subscribe", d.subscribe)
	routers.DELETE("/driver/unsubscribe", d.unSubscribe)

}
