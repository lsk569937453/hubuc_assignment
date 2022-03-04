package main

import (
	"fmt"
	"timer-go/controller"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

const PORT = "9394"

func main() {
	err := initController()
	if err != nil {
		glog.Fatalf("Initialization failed: %s", err)
		panic(err)
	}

}

func initController() error {
	r := gin.New()

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("[GIN] %v |%3d| %13v | %15s | %-7s  %#v %s |\"%s\" \n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
			param.Request.UserAgent(),
		)
	}))

	r.POST("/user", controller.RegisterUser)
	err := r.Run(":" + PORT) // listen and serve

	return err

}
