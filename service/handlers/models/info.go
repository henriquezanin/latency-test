package models

import (
	"TCC2/utils/logger"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
	"io/ioutil"
	"net/http"
)

func (p Params) Cpu() (string, *Error) {
	process, err := sysinfo.Self()
	if err != nil {
		modelErr := Error{
			InternalErr:      fmt.Errorf("error when create sysinfo: %v", err),
			InternalErrLevel: logger.Error,
			HttpStatusCode:   http.StatusInternalServerError,
		}
		return "", &modelErr
	}
	var cpu types.CPUTimer
	if castProcess, ok := process.(types.CPUTimer); !ok {
		return "", &Error{
			InternalErr:      fmt.Errorf("failed to get system info"),
			InternalErrLevel: logger.Error,
			HttpStatusCode:   http.StatusInternalServerError,
		}
	} else {
		cpu = castProcess
	}
	cpuTime, err := cpu.CPUTime()
	if err != nil {
		modelErr := Error{
			InternalErr:      fmt.Errorf("failed to get CPU info: %v", err),
			InternalErrLevel: logger.Error,
			HttpStatusCode:   http.StatusInternalServerError,
		}
		return "", &modelErr
	}
	b, err := json.Marshal(cpuTime)
	if err != nil {
		return "", &Error{
			InternalErr:      fmt.Errorf("failed to marshal json: %v", err),
			InternalErrLevel: logger.Error,
			HttpStatusCode:   http.StatusInternalServerError,
		}
	}
	data := string(b)
	infoChainData, err := p.cpuFromRequest("/cpu")
	if err != nil {
		return "", &Error{
			InternalErr:      fmt.Errorf("request data from service: %v", err),
			InternalErrLevel: logger.Error,
			HttpStatusCode:   http.StatusInternalServerError,
		}
	}
	return data + "\n" + infoChainData, nil
}

func (p Params) cpuFromRequest(path string) (string, error) {
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
