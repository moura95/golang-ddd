package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/moura95/go-ddd/internal/domain/vehicle"
	"go.uber.org/zap"
)

type vehicleRepository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}
type DtoVehicle struct {
	Uuid              uuid.UUID      `db:"uuid"`
	Brand             string         `db:"brand"`
	Model             string         `db:"model"`
	YearOfManufacture uint           `db:"year_of_manufacture"`
	LicensePlate      string         `db:"license_plate"`
	Color             string         `db:"color"`
	DeletedAt         sql.NullString `db:"deleted_at"`
	CreatedAt         time.Time      `db:"created_at"`
	UpdatedAt         time.Time      `db:"update_at"`
}

func NewVehicleRepository(db *sqlx.DB, log *zap.SugaredLogger) vehicle.IVehicleRepository {
	return &vehicleRepository{db: db, logger: log}
}

func (r *vehicleRepository) GetAll() ([]vehicle.Vehicle, error) {
	var vehicles []vehicle.Vehicle
	var dtoVehicles []DtoVehicle
	query := "SELECT * FROM vehicles WHERE deleted_at is null"
	if err := r.db.Select(&dtoVehicles, query); err != nil {
		return nil, err
	}
	for _, dto := range dtoVehicles {
		v := vehicle.NewVehicle(dto.Brand, dto.Model, dto.LicensePlate, dto.Color, dto.YearOfManufacture)
		vehicles = append(vehicles, *v)
	}

	return vehicles, nil
}

func (r *vehicleRepository) Create(ve vehicle.Vehicle) error {
	query := `
        INSERT INTO vehicles (brand, model, year_of_manufacture, license_plate, color)
        VALUES ($1, $2, $3, $4, $5)
    `
	args := []interface{}{
		ve.Brand,
		ve.Model,
		ve.YearOfManufacture,
		ve.LicensePlate,
		ve.Color,
	}
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
func (r *vehicleRepository) GetByID(uuid uuid.UUID) (*vehicle.Vehicle, error) {
	var dto DtoVehicle
	err := r.db.Get(&dto, "SELECT * FROM vehicles WHERE uuid = $1", uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Driver not found
		}
		return nil, err
	}
	v := vehicle.NewVehicle(dto.Brand, dto.Model, dto.LicensePlate, dto.Color, dto.YearOfManufacture)

	return v, nil
}

func (r *vehicleRepository) Update(uid uuid.UUID, ve *vehicle.Vehicle) error {
	query := `
        UPDATE vehicles 
        SET brand=$2, model=$3, year_of_manufacture=$4, license_plate=$5, color=$6, update_at=$7
        WHERE uuid=$1
    `
	args := []interface{}{
		uid,
		ve.Brand,
		ve.Model,
		ve.YearOfManufacture,
		ve.LicensePlate,
		ve.Color,
		time.Now(),
	}
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *vehicleRepository) HardDelete(uuid uuid.UUID) error {
	query := "DELETE FROM vehicles WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}
func (r *vehicleRepository) SoftDelete(uuid uuid.UUID) error {
	query := "UPDATE vehicles SET deleted_at=now() WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r *vehicleRepository) UnDelete(uuid uuid.UUID) error {
	query := "UPDATE vehicles SET deleted_at=null WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r *vehicleRepository) UnRelate(vehicleUUID uuid.UUID) error {
	query := "DELETE FROM drivers_vehicles WHERE vehicle_uuid = :VehicleUUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"VehicleUUID": vehicleUUID})
	return err
}
