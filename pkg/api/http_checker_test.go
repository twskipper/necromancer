/*
@Description: just go
@Author: skipper
@Date: 2019/11/14
@Time: 11:32 AM
@ProjectName necromancer
*/
package api

import (
	"log"
	"testing"
	"time"
)

func TestHTTPChecker_Run(t *testing.T) {
	hc := &HTTPChecker{
		Name:             "test",
		URL:              "http://example.com/",
		Method:           "GET",
		ExpectStatusCode: 200,
		Headers:          nil,
		Timeout:          5 * time.Second,
		Body:             nil,
	}
	res, err := hc.Run()
	if err != nil {
		t.Error(err)
	}
	log.Println("res: ", res)
}
