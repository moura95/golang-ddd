package vehicle

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/domain/vehicle"
	"github.com/moura95/go-ddd/internal/infra/cfg"
	"go.uber.org/zap"
)

type vehicleService struct {
	repository vehicle.IVehicleRepository
	config     cfg.Config
	logger     *zap.SugaredLogger
}

func NewVehicleService(repo vehicle.IVehicleRepository, cfg cfg.Config, log *zap.SugaredLogger) *vehicleService {
	return &vehicleService{
		repository: repo,
		config:     cfg,
		logger:     log,
	}
}

func (v *vehicleService) Create(brand, model, licensePlate, color string, yearOfManufacture uint) error {
	vehi := vehicle.NewVehicle(brand, model, licensePlate, color, yearOfManufacture)
	err := vehi.Validate()

	if err != nil {
		return err
	}
	err = v.repository.Create(*vehi)
	if err != nil {
		return fmt.Errorf("failed to create %s", err.Error())
	}
	return nil
}

func (v *vehicleService) List() ([]vehicle.Vehicle, error) {
	vehicles, err := v.repository.GetAll()
	if err != nil {
		return []vehicle.Vehicle{}, fmt.Errorf("failed to get %s", err.Error())

	}
	return vehicles, nil
}

func (v *vehicleService) GetByID(uid uuid.UUID) (vehicle.Vehicle, error) {
	ve, err := v.repository.GetByID(uid)
	if err != nil {
		return vehicle.Vehicle{}, fmt.Errorf("failed to get %s", err.Error())

	}
	return *ve, nil

}

func (v *vehicleService) Update(uid uuid.UUID, brand, model, licensePlate, color string, yearOfManufacture uint) error {
	vehi := vehicle.NewVehicle(brand, model, licensePlate, color, yearOfManufacture)

	err := v.repository.Update(uid, vehi)
	if err != nil {
		return fmt.Errorf("failed to update %s", err.Error())
	}
	return nil
}

func (v *vehicleService) SoftDelete(uid uuid.UUID) error {
	err := v.repository.SoftDelete(uid)
	if err != nil {
		return fmt.Errorf("failed to delete")
	}
	return nil
}
func (v *vehicleService) UnDelete(uid uuid.UUID) error {
	err := v.repository.UnDelete(uid)
	if err != nil {
		return fmt.Errorf("failed to un delete %s", err.Error())
	}
	return nil
}

func (v *vehicleService) HardDelete(uid uuid.UUID) error {
	// unRelate driver before delete
	err := v.repository.UnRelate(uid)
	if err != nil {
		return fmt.Errorf("failed to delete %s", err.Error())
	}

	err = v.repository.HardDelete(uid)
	if err != nil {
		return fmt.Errorf("failed to  delete %s", err.Error())
	}
	return nil
}
