package main

import (
	"town-guide/base"
	"town-guide/service"

	"github.com/kataras/iris/v12"
	irisrecover "github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	database := base.DbxDatabaseStarter{}
	database.Setup()

	app := iris.New()
	app.Use(irisrecover.New())
	userService := service.NewUserService()
	userRouter := app.Party("/user")
	userRouter.Get("/getUserInfo", userService.GetUserInfo)
	userRouter.Post("/register", userService.Register)

	//scenicService := service.NewScenicService()
	//groupRouter := app.Party("/scenic")
	//groupRouter.Post("/addScenic", scenicService)
	//groupRouter.Delete("/deleteScenic", scenicService.Register)
	//groupRouter.Put("/editScenic", a.Register)
	//groupRouter.Get("/getScenic", a.Register)
	//groupRouter.Get("/getAllScenic", a.Register)

	utilService := service.NewUtilService()
	groupRouter := app.Party("/util")
	groupRouter.Get("/getPic/{pic_name:string}", utilService.GetPic)
	groupRouter.Post("/uploadFile", utilService.UploadFile)
	groupRouter.Get("/getVideo/{video_name:string}", utilService.GetVideo)

	app.Run(iris.Addr(":8080"))
	//app.Run(iris.TLS(":443", "a.crt", "b.key"))
}
