package aggregate

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/domain/vehicle"
)

type DriverVehicleAggregate struct {
	Uuid          uuid.UUID
	Name          string
	Email         string
	TaxID         string
	DriverLicense string
	DateOfBirth   sql.NullString
	DeletedAt     sql.NullString
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Vehicles      []vehicle.Vehicle
}

type DriverVehicle struct {
	DriverUUID  uuid.UUID
	VehicleUUID uuid.UUID
}
