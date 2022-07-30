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
	userRouter.Get("/getUserDetail", userService.GetUserDetail)

	userActionRouter := app.Party("/user/action")
	userActionRouter.Post("/clock/{id}", userService.UserClockByScenicID)
	userActionRouter.Get("/checkScenicInfo", userService.Register)
	userActionRouter.Get("/wantCheckScenicInfo", userService.Register)

	scenicService := service.NewScenicService()
	groupRouter := app.Party("/scenic")
	groupRouter.Post("/addScenic", scenicService.AddScenic)
	groupRouter.Get("/getScenic/{scenic_id:int64}", scenicService.GetScenic)
	groupRouter.Delete("/deleteScenic/{scenic_id:int64}", scenicService.DeleteScenic)
	groupRouter.Get("/getAllScenic", scenicService.GetAllScenic)
	groupRouter.Post("/editScenic", scenicService.EditScenic)

	utilService := service.NewUtilService()
	groupRouter = app.Party("/utils")
	groupRouter.Get("/getPic/{pic_name:string}", utilService.GetPic)
	groupRouter.Post("/uploadFile", utilService.UploadFile)
	groupRouter.Get("/getVideo/{video_name:string}", utilService.GetVideo)

	app.Run(iris.Addr(":8080"))
	// app.Run(iris.TLS(":443", "a.crt", "b.key"))
}
