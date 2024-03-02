package vehicle

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestNewVehicle(t *testing.T) {
	vehicle := NewVehicle("Scania", "R500", "ABC123", "Blue", 2020)
	assert.Equal(t, "Scania", vehicle.Brand)
	assert.Equal(t, "R500", vehicle.Model)
	assert.Equal(t, "ABC123", vehicle.LicensePlate)
	assert.Equal(t, "Blue", vehicle.Color)
	err := vehicle.Validate()

	if err != nil {
		fmt.Println("Success TestNewVehicleWithOutBrand")
	}
}

func TestNewVehicleWithOutBrand(t *testing.T) {
	vehicle := NewVehicle("", "R500", "ABC123", "Blue", 2020)
	assert.Equal(t, "", vehicle.Brand)
	assert.Equal(t, "R500", vehicle.Model)
	assert.Equal(t, "ABC123", vehicle.LicensePlate)
	assert.Equal(t, "Blue", vehicle.Color)
	err := vehicle.Validate()
	if err.Error() == "invalid brand" {
		fmt.Println("Success TestNewVehicleWithOutBrand")
	}
}

func TestNewVehicleWithOutModel(t *testing.T) {
	vehicle := NewVehicle("Scania", "", "ABC123", "Blue", 2020)
	assert.Equal(t, "Scania", vehicle.Brand)
	assert.Equal(t, "", vehicle.Model)
	assert.Equal(t, "ABC123", vehicle.LicensePlate)
	assert.Equal(t, "Blue", vehicle.Color)
	err := vehicle.Validate()
	if err.Error() == "invalid model" {
		fmt.Println("Success TestNewVehicleWithOutModel")
	}
}

func TestNewVehicleWithOutLicensePlate(t *testing.T) {
	vehicle := NewVehicle("Scania", "R500", "", "Blue", 2020)
	assert.Equal(t, "Scania", vehicle.Brand)
	assert.Equal(t, "R500", vehicle.Model)
	assert.Equal(t, "", vehicle.LicensePlate)
	assert.Equal(t, "Blue", vehicle.Color)
	err := vehicle.Validate()
	if err.Error() == "invalid license plate" {
		fmt.Println("Success TestNewVehicleWithOutLicensePlate")
	}
}
