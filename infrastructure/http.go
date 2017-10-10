package infrastructure

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	HttpHandler interface {
		Get(string) ([]byte, error)
	}

	httpHandler struct {
		logger Logger
	}

	Logger interface {
		Debugf(string, ...interface{})
	}
)

func NewHttpHandler(logger Logger) HttpHandler {
	return &httpHandler{
		logger: logger,
	}
}

func (handler *httpHandler) Get(url string) ([]byte, error) {
	handler.logger.Debugf("HttpHandler Get url: %s", url)
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
