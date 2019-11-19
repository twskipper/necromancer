/*
@Description: just go
@Author: skipper
@Date: 2019/11/14
@Time: 11:21 AM
@ProjectName necromancer
*/
package api

type ContainerChecker struct {
	Name string
	Image string
	ExpectExitCode int
	Env map[string]string
}

func (c *ContainerChecker)Run()error  {
	return nil
}