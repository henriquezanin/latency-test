package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Error struct {
	InternalErr      error
	InternalErrLevel int
	HttpErr          string
	HttpStatusCode   int
}

type Params struct {
	ServiceChainURL string
}

func (p Params) chainRequest(path string) (string, error) {
	if p.ServiceChainURL == "" {
		return "", nil
	}
	response, err := http.Get(p.ServiceChainURL + path)
	if err != nil {
		return "", fmt.Errorf("failed to get http content: %v", err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}
	return string(responseData), nil
}
