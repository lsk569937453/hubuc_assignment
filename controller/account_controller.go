package controller

import (
	"net/http"
	"timer-go/constants"
	"timer-go/service"
	"timer-go/vojo"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func RegisterUser(c *gin.Context) {

	var res vojo.BaseRes
	res.Rescode = constants.NORMAL_RESPONSE_STATUS

	userId, err := service.RegisterUser(c)
	if err != nil {
		res.Rescode = constants.ERROR_RESPONSE_STATUS
		res.ResMessage = err.Error()
		glog.Errorf("register failed: %s", err)
	} else {
		res.ResMessage = userId
	}
	c.JSON(http.StatusOK, res)
}
