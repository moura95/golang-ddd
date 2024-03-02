package driver

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewDriver(t *testing.T) {
	driver := NewDriver("Motorista 1", "motorista1@example.com", "12345678901", "ABC12345", "1990-01-01")
	assert.Equal(t, "Motorista 1", driver.Name)
	assert.Equal(t, "motorista1@example.com", driver.Email)
	assert.Equal(t, "12345678901", driver.TaxID)
}

func TestNewDriverWithOutName(t *testing.T) {
	driver := NewDriver("", "motorista1@example.com", "12345678901", "ABC12345", "1990-01-01")
	assert.Equal(t, "", driver.Name)
	assert.Equal(t, "motorista1@example.com", driver.Email)
	assert.Equal(t, "12345678901", driver.TaxID)
}

func TestNewDriverWithOutTaxId(t *testing.T) {
	driver := NewDriver("Motorista 2", "motorista1@example.com", "", "ABC12345", "1990-01-01")
	assert.Equal(t, "Motorista 2", driver.Name)
	assert.Equal(t, "motorista1@example.com", driver.Email)
	assert.Equal(t, "", driver.TaxID)
}

func TestNewDriverWithOutEmail(t *testing.T) {
	driver := NewDriver("motorista 3", "", "12345678901", "ABC12345", "1990-01-01")
	assert.Equal(t, "motorista 3", driver.Name)
	assert.Equal(t, "", driver.Email)
	assert.Equal(t, "12345678901", driver.TaxID)

}
