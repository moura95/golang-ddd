package driver_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/domain/driver"
	"github.com/moura95/go-ddd/internal/infra/repository/memory"
	"github.com/stretchr/testify/assert"
)

type DriverServiceTest struct {
	repository driver.IDriverRepository
}

func NewDriverServiceTest() *DriverServiceTest {
	repo := memory.NewDriverRepository()
	return &DriverServiceTest{
		repository: repo,
	}
}

func TestCreateDriver(t *testing.T) {
	service := NewDriverServiceTest()

	d := driver.NewDriver("Driver Test", "test@example.com", "1234567890", "XYZ12345", "19/09/1995")

	err := service.repository.Create(*d)
	if err != nil {
		t.Error("Failed to created")
	}
	assert.NoError(t, err)

}

func TestGetAll(t *testing.T) {
	service := NewDriverServiceTest()

	drivers, err := service.repository.GetAll()
	if err != nil {
		t.Error("Failed to created")
	}
	assert.NoError(t, err)
	assert.Equal(t, drivers[0].Uuid, uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321"))
	assert.Equal(t, drivers[0].Name, "Driver 1")
	assert.Equal(t, drivers[0].Email, "driver1@example.com")
	assert.Equal(t, drivers[0].TaxID, "1234567890")
	assert.Equal(t, drivers[0].DriverLicense, "ABC12345")

	// assert record 2
	assert.Equal(t, drivers[1].Uuid, uuid.MustParse("ef9da75e-949f-4780-92b5-eda71618fc6c"))
	assert.Equal(t, drivers[1].Name, "Driver 2")
	assert.Equal(t, drivers[1].Email, "driver2@example.com")
	assert.Equal(t, drivers[1].TaxID, "9876543210")
	assert.Equal(t, drivers[1].DriverLicense, "XYZ98765")
}

func TestGetID(t *testing.T) {
	service := NewDriverServiceTest()

	dr, err := service.repository.GetByID(uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321"))
	if err != nil {
		t.Error("Failed to get")
	}
	assert.NoError(t, err)
	assert.Equal(t, dr.Uuid, uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321"))
	assert.Equal(t, dr.Name, "Driver 1")
	assert.Equal(t, dr.Email, "driver1@example.com")
	assert.Equal(t, dr.TaxID, "1234567890")
	assert.Equal(t, dr.DriverLicense, "ABC12345")

}

func TestUpdate(t *testing.T) {
	service := NewDriverServiceTest()

	uid := uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321")

	d := driver.NewDriver("Drive Updated", "driver12345@example.com", "123456", "ABC123451", "20/01/2000")

	err := service.repository.Update(uid, d)
	if err != nil {
		t.Error("Failed to update")
	}
	assert.NoError(t, err)
}

func TestHardDelete(t *testing.T) {
	service := NewDriverServiceTest()

	uid := uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321")
	err := service.repository.HardDelete(uid)
	if err != nil {
		t.Error("Failed to delete")
	}
	assert.NoError(t, err)
}

func TestSoftDelete(t *testing.T) {
	service := NewDriverServiceTest()

	uid := uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321")
	err := service.repository.SoftDelete(uid)
	if err != nil {
		t.Error("Failed to delete")
	}
	assert.NoError(t, err)
}
