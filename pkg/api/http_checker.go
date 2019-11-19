/*
@Description: just go
@Author: skipper
@Date: 2019/11/14
@Time: 11:00 AM
@ProjectName necromancer
*/
package api

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPChecker struct {
	Name             string
	URL              string
	Method           string
	ExpectStatusCode int
	Headers          []string
	Timeout          time.Duration
	Body             io.Reader
}

func (h *HTTPChecker) Run() (result *Result,err error) {
	now := time.Now()
	result = &Result{
		Name:     h.Name,
	}
	client := &http.Client{
		Timeout: h.Timeout,
	}
	req,err := http.NewRequest(h.Method,h.URL,h.Body)
	if err != nil {
		return nil,err
	}
	response,err := client.Do(req)
	if err != nil {
		return nil,err
	}
	if response.StatusCode != h.ExpectStatusCode {
		return nil,fmt.Errorf("not expected status code")
	}
	result.Ok = true
	result.Duration = time.Since(now)
	return result,nil
}
