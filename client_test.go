package coingate

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	client := NewClient()
	rate, _ := client.GetRate("USD", "CNY")
	fmt.Printf("rate: %.2f\n", rate)
}
