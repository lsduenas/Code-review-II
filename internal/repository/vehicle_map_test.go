package repository

import (
	"app/internal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVehicleMap_FindByColorAndYear(t *testing.T) {
	t.Run("success in retreving vehicles by color and year", func(t *testing.T) {
		// Arrange
		db := map[int]internal.Vehicle{
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
			2: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Mazda",
					Model:           "2024",
					Registration:    "ABC123",
					Color:           "red",
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
		repo := NewRepositoryReadVehicleMap(db)

		expectedResponse := map[int]internal.Vehicle{
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

		// Act
		obtainedResponse, err := repo.FindByColorAndYear("blue", 2023)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, obtainedResponse)
	})
	t.Run("failure in retreving vehicles by color and year", func(t *testing.T) {
		// Arrange
		db := map[int]internal.Vehicle{
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
			2: {
				Id: 1,
				VehicleAttributes: internal.VehicleAttributes{
					Brand:           "Mazda",
					Model:           "2024",
					Registration:    "ABC123",
					Color:           "red",
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
		repo := NewRepositoryReadVehicleMap(db)

		expectedResponse := map[int]internal.Vehicle{}

		// Act
		obtainedResponse, err := repo.FindByColorAndYear("orange", 2010)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, obtainedResponse)
	})
}
