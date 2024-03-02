package driver

import (
	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/aggregate"
)

type IDriverRepository interface {
	GetAll() ([]Driver, error)
	Create(driver Driver) error
	Subscribe(vehicle aggregate.DriverVehicle) error
	UnSubscribe(aggregate.DriverVehicle) error
	GetByID(uuid.UUID) (*aggregate.DriverVehicleAggregate, error)
	Update(uuid.UUID, *Driver) error
	HardDelete(uuid.UUID) error
	SoftDelete(uuid.UUID) error
	UnDelete(uuid.UUID) error
	UnRelate(uuid uuid.UUID) error
}
