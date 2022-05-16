package main

import (
	"fmt"

	coingate "github.com/ivanglie/go-coingate-client"
)

func main() {
	client := coingate.NewClient()
	rate, err := client.GetRate("USD", "CNY")
	if err != nil {
		panic(err)
	}
	fmt.Println(rate)
}
