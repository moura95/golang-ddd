package driver

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Driver struct {
	Uuid          uuid.UUID
	Name          string
	Email         string
	TaxID         string
	DriverLicense string
	DateOfBirth   sql.NullString
	DeletedAt     sql.NullString
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewDriver(name, email, taxId, driverLicense, dateOfBirth string) *Driver {
	return &Driver{
		Uuid:          uuid.New(),
		Name:          name,
		Email:         email,
		TaxID:         taxId,
		DriverLicense: driverLicense,
		DateOfBirth: sql.NullString{
			String: dateOfBirth,
			Valid:  false,
		},
	}
}

func (d *Driver) Validate() error {
	if d.Name == "" {
		return errors.New("invalid name")
	}
	if d.TaxID == "" {
		return errors.New("invalid tax id")
	}
	if d.Email == "" {
		return errors.New("invalid email")
	}
	return nil
}
