package client

import (
	"encoding/json"
	"github.com/evodicka/goccupied-client/io"
	"net/http"
	"strconv"
)

func GetValue() bool {
	config := io.ReadConfig()
	url := config.Http.Host + io.REQUESTPATH
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	responseBody := make(map[string]string)
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(responseBody)

	if value, err := strconv.ParseBool(responseBody["occupied"]); err != nil {
		return value
	} else {
		panic(err)
	}
}
