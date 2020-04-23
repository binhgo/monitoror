// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	configmodels "github.com/monitoror/monitoror/api/config/models"
	mock "github.com/stretchr/testify/mock"

	models "github.com/monitoror/monitoror/models"

	params "github.com/monitoror/monitoror/internal/pkg/monitorable/params"
)

// GeneratorEnabler is an autogenerated mock type for the GeneratorEnabler type
type GeneratorEnabler struct {
	mock.Mock
}

// Enable provides a mock function with given fields: variantName, generatorParamsValidator, tileGeneratorFunction
func (_m *GeneratorEnabler) Enable(variantName models.VariantName, generatorParamsValidator params.Validator, tileGeneratorFunction configmodels.TileGeneratorFunction) {
	_m.Called(variantName, generatorParamsValidator, tileGeneratorFunction)
}