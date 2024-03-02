package memory

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/aggregate"
	"github.com/moura95/go-ddd/internal/domain/driver"
)

type DriverRepositoryMemory struct {
	drivers       []driver.Driver
	driverVehicle []aggregate.DriverVehicle
}

func NewDriverRepository() *DriverRepositoryMemory {
	return &DriverRepositoryMemory{
		drivers: []driver.Driver{
			{
				Uuid:          uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321"),
				Name:          "Driver 1",
				Email:         "driver1@example.com",
				TaxID:         "1234567890",
				DriverLicense: "ABC12345",
				DateOfBirth:   sql.NullString{String: "1990-01-01", Valid: true},
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
			{
				Uuid:          uuid.MustParse("ef9da75e-949f-4780-92b5-eda71618fc6c"),
				Name:          "Driver 2",
				Email:         "driver2@example.com",
				TaxID:         "9876543210",
				DriverLicense: "XYZ98765",
				DateOfBirth:   sql.NullString{String: "1985-02-15", Valid: true},
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
		},
	}
}

func (m *DriverRepositoryMemory) GetAll() ([]driver.Driver, error) {

	return m.drivers, nil
}

func (m *DriverRepositoryMemory) Create(dr driver.Driver) error {
	driverInstance := driver.NewDriver(dr.Name, dr.Email, dr.TaxID, dr.DriverLicense, dr.DateOfBirth.String)
	m.drivers = append(m.drivers, *driverInstance)
	return nil
}

func (m *DriverRepositoryMemory) GetByID(u uuid.UUID) (*aggregate.DriverVehicleAggregate, error) {
	for _, d := range m.drivers {
		if d.Uuid == u {

			return &aggregate.DriverVehicleAggregate{
				Uuid:          d.Uuid,
				Name:          d.Name,
				Email:         d.Email,
				TaxID:         d.TaxID,
				DriverLicense: d.DriverLicense,
				DateOfBirth:   d.DateOfBirth,
				DeletedAt:     sql.NullString{},
				CreatedAt:     time.Time{},
				UpdatedAt:     time.Time{},
				Vehicles:      nil,
			}, nil
		}
	}
	return nil, nil
}

func (m *DriverRepositoryMemory) Update(u uuid.UUID, dr *driver.Driver) error {
	for i, d := range m.drivers {
		if d.Uuid == u {
			m.drivers[i] = *driver.NewDriver(dr.Name, dr.Email, dr.TaxID, dr.DriverLicense, dr.DateOfBirth.String)
			return nil
		}
	}
	return nil
}

func (m *DriverRepositoryMemory) HardDelete(u uuid.UUID) error {
	for i, dr := range m.drivers {
		if dr.Uuid == u {
			m.drivers = append(m.drivers[:i], m.drivers[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *DriverRepositoryMemory) SoftDelete(u uuid.UUID) error {
	for i, dr := range m.drivers {
		if dr.Uuid == u {
			dr.DeletedAt.String = time.Now().String()
			m.drivers[i] = dr
			return nil
		}
	}
	return nil
}

func (m *DriverRepositoryMemory) UnDelete(u uuid.UUID) error {
	for i, dr := range m.drivers {
		if dr.Uuid == u {
			dr.DeletedAt.String = ""
			m.drivers[i] = dr
			return nil
		}
	}
	return nil
}

func (m *DriverRepositoryMemory) Subscribe(ag aggregate.DriverVehicle) error {
	dvInstance := aggregate.DriverVehicle{
		DriverUUID:  ag.DriverUUID,
		VehicleUUID: ag.VehicleUUID,
	}
	m.driverVehicle = append(m.driverVehicle, dvInstance)
	return nil

}

func (m *DriverRepositoryMemory) UnSubscribe(ag aggregate.DriverVehicle) error {
	for i, dv := range m.driverVehicle {
		if dv.DriverUUID == ag.DriverUUID && dv.VehicleUUID == ag.DriverUUID {
			m.driverVehicle = append(m.driverVehicle[:i], m.driverVehicle[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *DriverRepositoryMemory) UnRelate(uuid.UUID) error {
	return nil
}
