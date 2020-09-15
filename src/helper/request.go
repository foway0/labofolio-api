package helper

import (
	"io/ioutil"
	"net/http"
)

// TODO helper & FOR TEST Write Happy Case
func Request(METHOD string, URL string) []byte {
	req, _ := http.NewRequest(METHOD, URL, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")

	client := new(http.Client)
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}