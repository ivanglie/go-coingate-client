package coingate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

const baseURL = "https://api.coingate.com/v2/rates/merchant"

type Error struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

func getCurrencyRate(from, to string, fetch FetchFunction) (float64, error) {
	log.Printf("Fetching the currency rate for %s\n", to)
	url := fmt.Sprintf("%s/%s/%s", baseURL, from, to)
	resp, err := fetch(url)
	if err != nil {
		return 0, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		var error Error
		err = json.Unmarshal(b, &error)
		if err != nil {
			return 0, err
		}
		return 0, fmt.Errorf("Service error (message: %s, reason: %s)", error.Message, error.Reason)
	}
	s := string(b)
	var res float64
	if res, err = strconv.ParseFloat(s, 64); err != nil {
		return 0, err
	}
	return res, nil
}
