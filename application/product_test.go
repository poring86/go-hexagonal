package application_test

import (
	"testing"

	"github.com/poring86/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestApplicationProduct_Enable(t *testing.T) {
	product := application.Product{}

	product.Name = "Aftershave"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be zero in order to have the product disabled", err.Error())
}
