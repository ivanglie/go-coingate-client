package coingate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	client := NewClient()
	rate, err := client.GetRate("USD", "CNY")
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, rate, float64(1))
}

func TestClientTheSameCurrency(t *testing.T) {
	client := NewClient()
	rate, err := client.GetRate("USD", "USD")
	assert.Nil(t, err)
	assert.Equal(t, float64(1), rate)
}

func TestClientWithError(t *testing.T) {
	client := NewClient()
	rate, err := client.GetRate("", "EUR")
	assert.Error(t, err)
	assert.Equal(t, float64(0), rate)
}
