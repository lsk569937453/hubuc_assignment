package service

import (
	"errors"
	"sync"
	"timer-go/vojo"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var UserMap sync.Map

func RegisterUser(c *gin.Context) (*string, error) {
	var req vojo.RegisterReq
	err := c.Bind(&req)

	if err != nil {
		return nil, err
	}
	if *req.Email == "" || *req.Password == "" || *req.UserName == "" {
		return nil, errors.New("field is empty")
	}

	if len(*req.Password) < 6 || len(*req.Password) > 120 {
		return nil, errors.New("password is illegal")
	}
	_, ok := UserMap.Load(*req.Email)
	if ok {
		return nil, errors.New("email is registered")
	}
	_, ok = UserMap.Load(*req.UserName)
	if ok {
		return nil, errors.New("userName is registered")
	}

	userId := uuid.NewV4().String()

	UserMap.Store(*req.Email, "")
	UserMap.Store(*req.UserName, "")

	return &userId, nil

}
