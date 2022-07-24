package controller

import (
	"town-guide/service"
)

// UserController
type UserController struct {
	Service service.UserService
}

func (c *UserController) GetBy(id string) service.Response {
	return c.Service.GetUserInfo(id)
}
