package handler

import (
	"app/internal"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVehicleService struct {
	mock.Mock
}

func (m *MockVehicleService) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(color, fabricationYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *MockVehicleService) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	args := m.Called(brand, startYear, endYear)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func (m *MockVehicleService) AverageMaxSpeedByBrand(brand string) (a float64, err error) {
	args := m.Called(brand)
	return args.Get(0).(float64), args.Error(1)
}

func (m *MockVehicleService) AverageCapacityByBrand(brand string) (a int, err error) {
	args := m.Called(brand)
	return args.Get(0).(int), args.Error(1)
}

func (m *MockVehicleService) SearchByWeightRange(query internal.SearchQuery, ok bool) (v map[int]internal.Vehicle, err error) {
	args := m.Called(query, ok)
	return args.Get(0).(map[int]internal.Vehicle), args.Error(1)
}

func TestHandlerVehicle_FindByColorAndYear(t *testing.T) {
	t.Run("success in retrieving vehicle by color and year", func(t *testing.T) {
		// Arrange
		mockService := new(MockVehicleService)

		handler := NewHandlerVehicle(mockService)
		r := chi.NewRouter()
		r.Get("/vehicles/{color}/{year}", handler.FindByColorAndYear())

		expectedVehicles := map[int]internal.Vehicle{
			1: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Toyota",
					Model:           "Corolla",
					Registration:    "ABC123",
					Color:           "blue",
					FabricationYear: 2023,
					Capacity:        5,
					MaxSpeed:        180.0,
					FuelType:        "Gasoline",
					Transmission:    "Automatic",
					Weight:          1500.0,
					Dimensions: internal.Dimensions{
						Height: 1.5,
						Length: 4.5,
						Width:  2.0,
					},
				},
			},
		}

		mockService.On("FindByColorAndYear", "blue", 2023).Return(expectedVehicles, nil)

		// Act 
		req := httptest.NewRequest("GET", "/vehicles/blue/2023", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		// Assert 
		assert.Equal(t, http.StatusOK, w.Code)

		// Verificar si se incluye la información correcta en la respuesta.
		expectedResponse := `
	{
		"data": {
			"1": {
				"Id": 1,
				"Brand": "Toyota",
				"Capacity": 5,
				"Color": "blue",
				"FabricationYear": 2023,
				"FuelType": "Gasoline",
				"Height": 1.5,
				"Length": 4.5,
				"MaxSpeed": 180.0,
				"Model": "Corolla",
				"Registration": "ABC123",
				"Transmission": "Automatic",
				"Weight": 1500.0,
				"Width": 2.0
			}
		},
		"message": "vehicles found"
	}
	`
		assert.JSONEq(t, expectedResponse, w.Body.String())
	})
}
