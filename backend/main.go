package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	irisrecover "github.com/kataras/iris/v12/middleware/recover"
	log "github.com/sirupsen/logrus"

	"time"
	"town-guide/base"
	"town-guide/service"
)

func main() {
	database := base.DbxDatabaseStarter{}
	database.Setup()

	app := iris.New()
	app.Use(irisrecover.New())
	cfg := logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              true,
		Columns:            true,
		MessageContextKeys: []string{"logger_message"},
		MessageHeaderKeys:  []string{"User-Agent"},
		LogFunc: func(now time.Time, latency time.Duration,
			status, ip, method, path string,
			message interface{},
			headerMessage interface{}) {
			app.Logger().Infof("| %s | %s | %s | %s | %s | %s | %+v | %+v",
				now.Format("2006-01-02.15:04:05.000000"),
				latency.String(), status, ip, method, path, headerMessage, message,
			)
		},
	}
	app.Use(logger.New(cfg))
	logger := app.Logger()
	logger.Install(log.StandardLogger())
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

	// app.Run(iris.Addr(":8080"))
	//app.Run(iris.TLS(":443", "a.crt", "b.key"))

}
