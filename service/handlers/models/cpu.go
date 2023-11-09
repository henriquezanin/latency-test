package models

import (
	"TCC2/utils/logger"
	"fmt"
	"net/http"
)

func (p Params) Cpu() (string, *Error) {
	data := "{user:359630000000,system:183070000000}"
	infoChainData, err := p.chainRequest("/cpu")
	if err != nil {
		return "", &Error{
			InternalErr:      fmt.Errorf("request data from service: %v", err),
			InternalErrLevel: logger.Error,
			HttpStatusCode:   http.StatusInternalServerError,
		}
	}
	return data + "\n" + infoChainData, nil
}
