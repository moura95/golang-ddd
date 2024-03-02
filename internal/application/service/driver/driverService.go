package driverservice

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/aggregate"
	"github.com/moura95/go-ddd/internal/domain/driver"
	"github.com/moura95/go-ddd/internal/infra/cfg"
	"go.uber.org/zap"
)

type Service struct {
	repository driver.IDriverRepository
	config     cfg.Config
	logger     *zap.SugaredLogger
}

func NewDriverService(repo driver.IDriverRepository, cfg cfg.Config, log *zap.SugaredLogger) *Service {
	return &Service{
		repository: repo,
		config:     cfg,
		logger:     log,
	}
}

func (d *Service) Create(name, email, taxId, driverLicense, dateOfBirth string) error {
	dr := driver.NewDriver(name, email, taxId, driverLicense, dateOfBirth)
	err := dr.Validate()
	if err != nil {
		return err
	}

	err = d.repository.Create(*dr)
	if err != nil {
		return fmt.Errorf("failed to create %s", err.Error())
	}
	return nil
}
func (d *Service) Subscribe(driverUUID, vehicleUUID uuid.UUID) error {
	driverVehicle := aggregate.DriverVehicle{
		DriverUUID:  driverUUID,
		VehicleUUID: vehicleUUID,
	}
	err := d.repository.Subscribe(driverVehicle)
	if err != nil {
		return fmt.Errorf("failed to create %s", err.Error())
	}
	return nil
}
func (d *Service) UnSubscribe(driverUUID, vehicleUUID uuid.UUID) error {
	driverVehicle := aggregate.DriverVehicle{
		DriverUUID:  driverUUID,
		VehicleUUID: vehicleUUID,
	}
	err := d.repository.UnSubscribe(driverVehicle)
	if err != nil {
		return fmt.Errorf("failed to hard delete driver %s", err.Error())
	}
	return nil
}

func (d *Service) GetByID(uid uuid.UUID) (*aggregate.DriverVehicleAggregate, error) {
	driverOutput, err := d.repository.GetByID(uid)

	if err != nil {
		return nil, fmt.Errorf("failed to get driver %s", err.Error())
	}
	if driverOutput == nil {
		return nil, fmt.Errorf("not found")

	}

	return driverOutput, nil
}

func (d *Service) List() ([]driver.Driver, error) {
	drivers, err := d.repository.GetAll()
	if err != nil {
		return []driver.Driver{}, fmt.Errorf("failed to list drivers %s", err.Error())
	}
	return drivers, nil
}

func (d *Service) Update(uid uuid.UUID, name, email, taxId, driverLicense, dateOfBirth string) error {
	dr := driver.NewDriver(name, email, taxId, driverLicense, dateOfBirth)
	err := d.repository.Update(uid, dr)
	if err != nil {
		return fmt.Errorf("failed to update driver %s", err.Error())
	}
	return nil
}

func (d *Service) SoftDelete(uid uuid.UUID) error {
	err := d.repository.SoftDelete(uid)
	if err != nil {
		return fmt.Errorf("failed to delete driver %s", err.Error())
	}
	return nil
}

func (d *Service) UnDelete(uid uuid.UUID) error {
	err := d.repository.UnDelete(uid)
	if err != nil {
		return fmt.Errorf("failed to recover driver %s", err.Error())
	}
	return nil
}

func (d *Service) HardDelete(uid uuid.UUID) error {
	// unRelate driver before delete
	err := d.repository.UnRelate(uid)
	if err != nil {
		return fmt.Errorf("failed to hard delete driver %s", err.Error())
	}

	err = d.repository.HardDelete(uid)
	if err != nil {
		return fmt.Errorf("failed to hardelete driver %s", err.Error())
	}
	return nil
}
