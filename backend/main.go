package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"town-guide/controller"
	"town-guide/service"
)

func main() {
	app := iris.New()
	mvc.Configure(app.Party("/user"), userController)
	app.Run(iris.Addr(":8080"))
}

func userController(app *mvc.Application) {
	app.Handle(&controller.UserController{Service:
	service.NewUserService()})
}
