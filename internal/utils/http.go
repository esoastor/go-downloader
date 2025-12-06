package utils

import (
	"io"
	"net/http"
)

func MakeGetRequest(url string) []byte {
	response, error := http.Get(url)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	
	if (response.StatusCode != 200) {
		panic("Bad response")
	}

	body, error := io.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}
	return body
}

