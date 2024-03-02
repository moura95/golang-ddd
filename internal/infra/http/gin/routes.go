package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moura95/go-ddd/internal/application/service/driver"
	vehicleservice "github.com/moura95/go-ddd/internal/application/service/vehicle"
	driverrouter "github.com/moura95/go-ddd/internal/infra/http/gin/router/driver"
	vehiclerouter "github.com/moura95/go-ddd/internal/infra/http/gin/router/vehicle"
	driverpostgres "github.com/moura95/go-ddd/internal/infra/repository/postgres"
	"go.uber.org/zap"
)

func (s *Server) createRoutesV1(router *gin.Engine, log *zap.SugaredLogger) {
	router.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	routes := router.Group("/")
	// Instance Driver Repository Postgres
	driverRepository := driverpostgres.NewDriverRepository(s.store, log)
	// Instance Driver Service with Postgres
	driverService := driverservice.NewDriverService(driverRepository, *s.config, log)

	// Instance VehicleRepository Repository
	vehicleRepository := driverpostgres.NewVehicleRepository(s.store, log)
	// Instance VehicleRouter Service with Postgres
	vehicleService := vehicleservice.NewVehicleService(vehicleRepository, *s.config, log)

	vehiclerouter.NewVehicleRouter(vehicleService, log).SetupVehicleRoute(routes)
	driverrouter.NewDriverRouter(driverService, log).SetupDriverRoute(routes)
}
