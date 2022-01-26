package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func HttpPost(url string, body []byte) ([]byte, error) {
	res, err := http.Post(url,
		"application/json;charset=utf-8",
		bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}
