package models

import (
	"fmt"
	"github.com/Project-Retina/utils/logger"
	"net/http"
)

func (p Params) Text() (string, *Error) {
	requestText, err := p.chainRequest("/text")
	if err != nil {
		return "", &Error{
			InternalErr:      fmt.Errorf("request data from service: %v", err),
			InternalErrLevel: logger.Error,
			HttpStatusCode:   http.StatusInternalServerError,
		}
	}
	return loremIpsum + "\n" + requestText, nil
}
