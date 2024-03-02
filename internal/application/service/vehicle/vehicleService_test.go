package vehicle

import (
	"testing"

	"github.com/google/uuid"
	"github.com/moura95/go-ddd/internal/domain/vehicle"
	"github.com/moura95/go-ddd/internal/infra/repository/memory"
	"github.com/stretchr/testify/assert"
)

type VehicleServiceTest struct {
	repository memory.IVehicleRepositoryMemory
}

func NewVehicleServiceTest(repo memory.IVehicleRepositoryMemory) *VehicleServiceTest {
	return &VehicleServiceTest{
		repository: repo,
	}
}

func TestCreateVehicle(t *testing.T) {
	mockRepo := memory.NewVehicleRepository()
	service := NewVehicleServiceTest(mockRepo)

	ve := vehicle.Vehicle{
		Uuid:              uuid.New(),
		Brand:             "Volvo",
		Model:             "FJ15",
		YearOfManufacture: 2021,
		LicensePlate:      "FIT1585",
		Color:             "Blue",
	}

	err := service.repository.Create(ve)
	if err != nil {
		t.Error("Failed to created")
	}
	assert.NoError(t, err)

}

func TestGetAllVehicles(t *testing.T) {
	mockRepo := memory.NewVehicleRepository()

	service := NewVehicleServiceTest(mockRepo)

	vehicles, err := service.repository.GetAll()
	if err != nil {
		t.Error("Failed to created")
	}
	assert.NoError(t, err)
	assert.Equal(t, vehicles[0].Uuid, uuid.MustParse("43ee3d4c-de06-4021-ab6f-ba8113418df9"))
	assert.Equal(t, vehicles[0].Brand, "Scania")
	assert.Equal(t, vehicles[0].Model, "R500")
	assert.Equal(t, vehicles[0].YearOfManufacture, uint(2020))
	assert.Equal(t, vehicles[0].LicensePlate, "ABC123")
	assert.Equal(t, vehicles[0].Color, "Blue")

	// assert record 2
	assert.Equal(t, vehicles[1].Uuid, uuid.MustParse("457a8df2-782f-4f22-8233-623b694096a1"))
	assert.Equal(t, vehicles[1].Brand, "Volvo")
	assert.Equal(t, vehicles[1].Model, "FH16")
	assert.Equal(t, vehicles[1].YearOfManufacture, uint(2019))
	assert.Equal(t, vehicles[1].LicensePlate, "XYZ987")
	assert.Equal(t, vehicles[1].Color, "Black")
}

func TestGetVehicleID(t *testing.T) {
	mockRepo := memory.NewVehicleRepository()
	service := NewVehicleServiceTest(mockRepo)

	vehicle, err := service.repository.GetByID(uuid.MustParse("43ee3d4c-de06-4021-ab6f-ba8113418df9"))
	if err != nil {
		t.Error("Failed to get")
	}

	assert.NoError(t, err)
	assert.Equal(t, vehicle.Uuid, uuid.MustParse("43ee3d4c-de06-4021-ab6f-ba8113418df9"))
	assert.Equal(t, vehicle.Brand, "Scania")
	assert.Equal(t, vehicle.Model, "R500")
	assert.Equal(t, vehicle.YearOfManufacture, uint(2020))
	assert.Equal(t, vehicle.LicensePlate, "ABC123")
	assert.Equal(t, vehicle.Color, "Blue")

}

func TestUpdateVehicle(t *testing.T) {
	mockRepo := memory.NewVehicleRepository()
	service := NewVehicleServiceTest(mockRepo)

	uid := uuid.MustParse("43ee3d4c-de06-4021-ab6f-ba8113418df9")

	d := &vehicle.Vehicle{
		Uuid:              uid,
		Brand:             "Scania Update",
		Model:             "R501",
		YearOfManufacture: 2021,
		LicensePlate:      "ABC321",
		Color:             "RED",
	}

	err := service.repository.Update(d)
	if err != nil {
		t.Error("Failed to update")
	}
	assert.NoError(t, err)
}

func TestVehicleHardDelete(t *testing.T) {
	mockRepo := memory.NewVehicleRepository()
	service := NewVehicleServiceTest(mockRepo)

	uid := uuid.MustParse("43ee3d4c-de06-4021-ab6f-ba8113418df9")
	err := service.repository.HardDelete(uid)
	if err != nil {
		t.Error("Failed to delete")
	}
	assert.NoError(t, err)
}

func TestVehicleSoftDelete(t *testing.T) {
	mockRepo := memory.NewVehicleRepository()
	service := NewVehicleServiceTest(mockRepo)

	uid := uuid.MustParse("43ee3d4c-de06-4021-ab6f-ba8113418df9")
	err := service.repository.SoftDelete(uid)
	if err != nil {
		t.Error("Failed to delete")
	}
	assert.NoError(t, err)
}
