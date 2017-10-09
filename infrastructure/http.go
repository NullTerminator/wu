package infrastructure

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	HttpHandler struct {
	}
)

func (handler *HttpHandler) Get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Bad HTTP Status: %d\n", res.StatusCode))
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return body, nil
}
