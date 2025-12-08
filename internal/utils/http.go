package utils

import (
	"io"
	"net/http"
)

func MakeGetRequest(url string, callback func(resp *http.Response)) []byte {
    client := &http.Client{
        Transport: &http.Transport{},
    }
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        panic(err)
    }

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
	response, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	
	if (response.StatusCode != 200) {
		callback(response)	
	}

	body, error := io.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}
	return body
}

