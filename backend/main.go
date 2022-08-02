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
	userRouter.Get("/getUserDetail/{open_id:string}", userService.GetUserDetail)

	userActionRouter := app.Party("/user/action")
	userActionRouter.Post("/clock", userService.UserClockByScenicID)
	userActionRouter.Post("/collection", userService.UserCollectionByScenicID)

	scenicService := service.NewScenicService()
	groupRouter := app.Party("/scenic")
	groupRouter.Post("/add", scenicService.AddScenic)
	groupRouter.Get("/get/{scenic_id:int64}", scenicService.GetScenic)
	groupRouter.Delete("/delete/{scenic_id:int64}", scenicService.DeleteScenic)
	groupRouter.Get("/getAll", scenicService.GetAllScenic)
	groupRouter.Get("/getByCategory/{category_id:int64}", scenicService.GetScenicByCategoryID)
	groupRouter.Post("/edit", scenicService.EditScenic)

	categoryService := service.NewCategoryService()
	categoryRouter := app.Party("/category")
	categoryRouter.Post("/add", categoryService.AddCategory)
	categoryRouter.Delete("/delete/{category_id:int64}", categoryService.DeleteCategory)
	categoryRouter.Get("/getAll", categoryService.GetAllCategory)
	categoryRouter.Post("/edit", categoryService.EditCategory)

	utilService := service.NewUtilService()
	groupRouter = app.Party("/utils")
	groupRouter.Get("/getPic/{pic_name:string}", utilService.GetPic)
	groupRouter.Post("/uploadFile", utilService.UploadFile)
	groupRouter.Get("/getMedia/{media_name:string}", utilService.GetMedia)

///	app.Run(iris.Addr(":8080"))
	app.Run(iris.TLS(":443", "a.crt", "b.key"))
}
