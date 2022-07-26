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
	app.Run(iris.TLS("127.0.0.1:443","a.crt","b.key"))
}

func userController(app *mvc.Application) {
	app.Handle(&controller.UserController{Service:
	service.NewUserService()})
}
