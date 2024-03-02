package vehicle

import (
	"github.com/google/uuid"
)

type IVehicleRepository interface {
	GetAll() ([]Vehicle, error)
	Create(Vehicle) error
	GetByID(uuid.UUID) (*Vehicle, error)
	Update(uuid.UUID, *Vehicle) error
	HardDelete(uuid.UUID) error
	SoftDelete(uuid.UUID) error
	UnDelete(uuid.UUID) error
	UnRelate(uuid.UUID) error
}
