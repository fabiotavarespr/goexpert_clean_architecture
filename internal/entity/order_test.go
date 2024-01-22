package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "e1a4a91e-aa6d-4fe3-849a-59affc25bbdf"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "e1a4a91e-aa6d-4fe3-849a-59affc25bbdf", Price: 10}
	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenICallNewOrder_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order := Order{
		ID:    "e1a4a91e-aa6d-4fe3-849a-59affc25bbdf",
		Price: 10.0,
		Tax:   2.0,
	}
	assert.Equal(t, "e1a4a91e-aa6d-4fe3-849a-59affc25bbdf", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.IsValid())
}

func TestGivenAValidParams_WhenICallNewOrderFunc_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("e1a4a91e-aa6d-4fe3-849a-59affc25bbdf", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Equal(t, "e1a4a91e-aa6d-4fe3-849a-59affc25bbdf", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
}

func TestGivenAPriceAndTax_WhenICallCalculatePrice_ThenIShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("e1a4a91e-aa6d-4fe3-849a-59affc25bbdf", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	assert.Equal(t, 12.0, order.FinalPrice)
}
