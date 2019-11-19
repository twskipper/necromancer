/*
@Description: just go
@Author: skipper
@Date: 2019/11/14
@Time: 11:13 AM
@ProjectName necromancer
*/
package api

import "time"

type Checker interface {
	Run() (*Result, error)
}

type Result struct {
	Name     string
	Id       string
	msg      string
	Ok       bool
	Duration time.Duration
}
