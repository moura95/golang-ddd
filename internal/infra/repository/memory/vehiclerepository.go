package memory

import (
	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/domain/vehicle"

	"time"
)

type VehicleRepositoryMemory struct {
	vehicles []vehicle.Vehicle
}

func NewVehicleRepository() *VehicleRepositoryMemory {
	return &VehicleRepositoryMemory{vehicles: []vehicle.Vehicle{
		{
			Uuid:              uuid.MustParse("43ee3d4c-de06-4021-ab6f-ba8113418df9"),
			Brand:             "Scania",
			Model:             "R500",
			YearOfManufacture: uint(2020),
			LicensePlate:      "ABC123",
			Color:             "Blue",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
		{
			Uuid:              uuid.MustParse("457a8df2-782f-4f22-8233-623b694096a1"),
			Brand:             "Volvo",
			Model:             "FH16",
			YearOfManufacture: uint(2019),
			LicensePlate:      "XYZ987",
			Color:             "Black",
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		}}}
}

func (v VehicleRepositoryMemory) GetAll() ([]vehicle.Vehicle, error) {
	return v.vehicles, nil
}

func (v VehicleRepositoryMemory) Create(ve vehicle.Vehicle) error {
	v.vehicles = append(v.vehicles, ve)
	return nil
}

func (v VehicleRepositoryMemory) GetByID(u uuid.UUID) (*vehicle.Vehicle, error) {
	for _, ve := range v.vehicles {
		if ve.Uuid == u {
			return &ve, nil
		}
	}
	return nil, nil
}

func (v VehicleRepositoryMemory) Update(uid uuid.UUID, vehi *vehicle.Vehicle) error {
	for i, ve := range v.vehicles {
		if ve.Uuid == uid {
			v.vehicles[i] = *vehi
			return nil
		}
	}
	return nil
}

func (v VehicleRepositoryMemory) HardDelete(u uuid.UUID) error {
	for i, ve := range v.vehicles {
		if ve.Uuid == u {
			v.vehicles = append(v.vehicles[:i], v.vehicles[i+1:]...)
			return nil
		}
	}
	return nil
}

func (v VehicleRepositoryMemory) SoftDelete(u uuid.UUID) error {
	for i, ve := range v.vehicles {
		if ve.Uuid == u {
			ve.DeletedAt.String = time.Now().String()
			v.vehicles[i] = ve
			return nil
		}
	}
	return nil
}

func (v VehicleRepositoryMemory) UnDelete(u uuid.UUID) error {
	for i, ve := range v.vehicles {
		if ve.Uuid == u {
			ve.DeletedAt.String = ""
			v.vehicles[i] = ve
			return nil
		}
	}
	return nil
}

func (v VehicleRepositoryMemory) UnRelate(uuid.UUID) error {
	return nil
}
