package vehicle

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Vehicle struct {
	Uuid              uuid.UUID
	Brand             string
	Model             string
	YearOfManufacture uint
	LicensePlate      string
	Color             string
	DeletedAt         sql.NullString
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func NewVehicle(brand, model, licensePlate, color string, yearOfManufacture uint) *Vehicle {
	return &Vehicle{
		Uuid:              uuid.New(),
		Brand:             brand,
		Model:             model,
		YearOfManufacture: yearOfManufacture,
		LicensePlate:      licensePlate,
		Color:             color,
	}
}

func (v *Vehicle) Validate() error {
	if v.Brand == "" {
		return errors.New("invalid brand")
	}
	if v.Model == "" {
		return errors.New("invalid model")
	}
	if v.LicensePlate == "" {
		return errors.New("invalid license plate")
	}
	return nil

}
