package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/moura95/go-ddd/internal/aggregate"
	"github.com/moura95/go-ddd/internal/domain/driver"
	"github.com/moura95/go-ddd/internal/domain/vehicle"
	"go.uber.org/zap"
)

type driverRepository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

type DriverVehicleModel struct {
	DriverUUID    uuid.UUID      `db:"driver_uuid"`
	DriverName    string         `db:"name"`
	DriverEmail   string         `db:"email"`
	DriverTaxID   string         `db:"tax_id"`
	DriverLicense string         `db:"driver_license"`
	DriverDOB     sql.NullString `db:"date_of_birth"`
	VehicleUUID   uuid.UUID      `db:"uuid"`
	VehicleBrand  string         `db:"brand"`
	VehicleModel  string         `db:"model"`
	VehicleYear   uint           `db:"year_of_manufacture"`
	VehicleColor  string         `db:"color"`
	CreatedAt     time.Time      `db:"created_at"`
	UpdatedAt     time.Time      `db:"update_at"`
}

type DriverModel struct {
	Uuid          uuid.UUID      `db:"uuid"`
	Name          string         `db:"name"`
	Email         string         `db:"email"`
	TaxID         string         `db:"tax_id"`
	DriverLicense string         `db:"driver_license"`
	DateOfBirth   sql.NullString `db:"date_of_birth"`
	DeletedAt     sql.NullString `db:"deleted_at"`
	CreatedAt     time.Time      `db:"created_at"`
	UpdatedAt     time.Time      `db:"update_at"`
}

func NewDriverRepository(db *sqlx.DB, log *zap.SugaredLogger) driver.IDriverRepository {
	return &driverRepository{db: db, logger: log}
}

func (r *driverRepository) GetAll() ([]driver.Driver, error) {
	var drivers []driver.Driver
	var model []DriverModel
	query := "SELECT * FROM drivers WHERE deleted_at is null"
	if err := r.db.Select(&model, query); err != nil {
		return []driver.Driver{}, err
	}
	for _, d := range model {
		instanceDriver := driver.NewDriver(d.Name, d.Email, d.TaxID, d.DriverLicense, d.DateOfBirth.String)
		instanceDriver.Uuid = d.Uuid
		drivers = append(drivers, *instanceDriver)
	}
	return drivers, nil
}

func (r *driverRepository) Create(dr driver.Driver) error {
	query := `
        INSERT INTO drivers (name, email, tax_id, driver_license, date_of_birth)
        VALUES ($1, $2, $3, $4, $5)
    `
	args := []interface{}{
		dr.Name,
		dr.Email,
		dr.TaxID,
		dr.DriverLicense,
		dr.DateOfBirth,
	}
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *driverRepository) Subscribe(driverVehicle aggregate.DriverVehicle) error {
	query := `
        INSERT INTO drivers_vehicles (driver_uuid, vehicle_uuid)
        VALUES ($1, $2)
    `
	args := []interface{}{
		driverVehicle.VehicleUUID,
		driverVehicle.DriverUUID,
	}
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *driverRepository) UnSubscribe(driverVehicle aggregate.DriverVehicle) error {
	query := "DELETE FROM drivers_vehicles WHERE driver_uuid =$1 AND vehicle_uuid =$2"
	args := []interface{}{
		driverVehicle.VehicleUUID,
		driverVehicle.DriverUUID,
	}
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return err
}

func (r *driverRepository) GetByID(uid uuid.UUID) (*aggregate.DriverVehicleAggregate, error) {
	var model []DriverVehicleModel

	query := `
		SELECT d.uuid driver_uuid, d.name, d.email, d.tax_id, d.driver_license, d.date_of_birth,
		       v.uuid, v.brand , v.model,
		       v.year_of_manufacture , v.color
		FROM drivers AS d
		LEFT JOIN drivers_vehicles AS dv ON d.uuid = dv.driver_uuid
		LEFT JOIN vehicles AS v ON v.uuid = dv.vehicle_uuid
		WHERE d.uuid = $1
	`

	err := r.db.Select(&model, query, uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || err.Error() == "missing destination name uuid in *[]postgres.DriverVehicleDb" {
			return nil, nil
		}
		return nil, err
	}
	if model == nil {
		return nil, nil
	}

	d := &aggregate.DriverVehicleAggregate{
		Uuid:          model[0].DriverUUID,
		Name:          model[0].DriverName,
		Email:         model[0].DriverEmail,
		TaxID:         model[0].DriverTaxID,
		DriverLicense: model[0].DriverLicense,
		DateOfBirth:   model[0].DriverDOB,
		Vehicles:      make([]vehicle.Vehicle, 0, len(model)),
	}

	for _, res := range model {
		v := vehicle.Vehicle{
			Uuid:              res.VehicleUUID,
			Brand:             res.VehicleBrand,
			Model:             res.VehicleModel,
			YearOfManufacture: res.VehicleYear,
			Color:             res.VehicleColor,
		}
		d.Vehicles = append(d.Vehicles, v)
	}

	return d, nil
}

func (r *driverRepository) Update(uuid uuid.UUID, dr *driver.Driver) error {
	query := `
        UPDATE drivers 
        SET name=$2, tax_id=$3, driver_license=$4, date_of_birth=$5, update_at=$6
    	WHERE uuid= $1`

	args := []interface{}{
		uuid,
		dr.Name,
		dr.TaxID,
		dr.DriverLicense,
		dr.DateOfBirth,
		time.Now(),
	}
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *driverRepository) HardDelete(uuid uuid.UUID) error {
	query := "DELETE FROM drivers WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r *driverRepository) SoftDelete(uuid uuid.UUID) error {
	query := "UPDATE drivers SET deleted_at=now() WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r *driverRepository) UnDelete(uuid uuid.UUID) error {
	query := "UPDATE drivers SET deleted_at=null WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r *driverRepository) UnRelate(driverUUID uuid.UUID) error {
	query := "DELETE FROM drivers_vehicles WHERE driver_uuid = :DriverUUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"DriverUUID": driverUUID})
	return err
}
