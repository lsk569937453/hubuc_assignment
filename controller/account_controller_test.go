package controller

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
	"timer-go/constants"
	"timer-go/util"
	"timer-go/vojo"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestAccountRegisterError1(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)

	urlIndex := "/user"

	userName := ""
	email := ""
	password := ""

	body := vojo.RegisterReq{
		UserName: &userName,
		Email:    &email,
		Password: &password,
	}
	headers := map[string]string{}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(400, w.Code)
}
func TestAccountRegisterError2(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)

	urlIndex := "/user"

	userName := ""
	email := "569937453@gmail.com"
	password := "1223315465"

	body := vojo.RegisterReq{
		UserName: &userName,
		Email:    &email,
		Password: &password,
	}
	headers := map[string]string{}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes
	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(nil, err)
	assert.Equal(constants.ERROR_RESPONSE_STATUS, baseRes.Rescode)

}
func TestAccountRegisterError3(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)

	urlIndex := "/user"

	userName := "lsk"
	email := ""
	password := "1223315465"

	body := vojo.RegisterReq{
		UserName: &userName,
		Email:    &email,
		Password: &password,
	}
	headers := map[string]string{}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(400, w.Code)

}
func TestAccountRegisterError4(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)

	urlIndex := "/user"

	userName := "lsk"
	email := "lsk569937453@gmail.com"
	password := ""

	body := vojo.RegisterReq{
		UserName: &userName,
		Email:    &email,
		Password: &password,
	}
	headers := map[string]string{}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes
	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(nil, err)
	assert.Equal(constants.ERROR_RESPONSE_STATUS, baseRes.Rescode)

}
func TestAccountRegisterError5(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)

	urlIndex := "/user"

	userName := "lsk"
	email := "xxxx."
	password := "1234567"

	body := vojo.RegisterReq{
		UserName: &userName,
		Email:    &email,
		Password: &password,
	}
	headers := map[string]string{}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(400, w.Code)

}
func TestAccountRegisterError6(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)

	urlIndex := "/user"

	userName := "lsk"
	email := "lsk@gmail.com."
	password := "123"

	body := vojo.RegisterReq{
		UserName: &userName,
		Email:    &email,
		Password: &password,
	}
	headers := map[string]string{}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes

	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(nil, err)
	assert.Equal(constants.ERROR_RESPONSE_STATUS, baseRes.Rescode)
	resMessageStr := baseRes.ResMessage.(string)
	assert.Equal("password is illegal", resMessageStr)
}
func TestAccountRegisterError7(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)

	urlIndex := "/user"

	userName := "lsk"
	email := "lsk@gmail.com."
	password := "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901"

	body := vojo.RegisterReq{
		UserName: &userName,
		Email:    &email,
		Password: &password,
	}
	headers := map[string]string{}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes

	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(nil, err)
	assert.Equal(constants.ERROR_RESPONSE_STATUS, baseRes.Rescode)
	resMessageStr := baseRes.ResMessage.(string)
	assert.Equal("password is illegal", resMessageStr)
}
func TestAccountRegisterError8(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)

	urlIndex := "/user"

	userName := "lsk8"
	email := "lsk5699374538@gmail.com"
	password := "123456"

	body := vojo.RegisterReq{
		UserName: &userName,
		Email:    &email,
		Password: &password,
	}
	headers := map[string]string{}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes

	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(nil, err)
	assert.Equal(0, baseRes.Rescode)

	userName2 := "lskplus"
	email2 := "lsk5699374538@gmail.com"
	password2 := "123456"

	body2 := vojo.RegisterReq{
		UserName: &userName2,
		Email:    &email2,
		Password: &password2,
	}
	//second register
	w = util.PostJson(urlIndex, body2, router, headers)
	assert.Equal(200, w.Code)
	resBody = w.Body.String()

	var newBaseRes vojo.BaseRes

	err = json.Unmarshal([]byte(resBody), &newBaseRes)
	assert.Equal(nil, err)
	assert.Equal(constants.ERROR_RESPONSE_STATUS, newBaseRes.Rescode)
	resMessageStr := newBaseRes.ResMessage.(string)
	assert.Equal("email is registered", resMessageStr)

}
func TestAccountRegisterError9(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)

	urlIndex := "/user"

	userName := "lsk9"
	email := "lsk5699374539@gmail.com"
	password := "123456"

	body := vojo.RegisterReq{
		UserName: &userName,
		Email:    &email,
		Password: &password,
	}
	headers := map[string]string{}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes

	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(nil, err)
	assert.Equal(0, baseRes.Rescode)

	userName2 := "lsk9"
	email2 := "lsk569937453plus@gmail.com"
	password2 := "123456"

	body2 := vojo.RegisterReq{
		UserName: &userName2,
		Email:    &email2,
		Password: &password2,
	}
	//second register
	w = util.PostJson(urlIndex, body2, router, headers)
	assert.Equal(200, w.Code)
	resBody = w.Body.String()

	var newBaseRes vojo.BaseRes

	err = json.Unmarshal([]byte(resBody), &newBaseRes)
	assert.Equal(nil, err)
	assert.Equal(constants.ERROR_RESPONSE_STATUS, newBaseRes.Rescode)
	resMessageStr := newBaseRes.ResMessage.(string)
	assert.Equal("userName is registered", resMessageStr)

}
func TestAccountRegisterOk1(t *testing.T) {
	var w *httptest.ResponseRecorder
	assert := assert.New(t)

	urlIndex := "/user"

	userName := "lsk10"
	email := "lsk56993745310@gmail.com."
	password := "123456"

	body := vojo.RegisterReq{
		UserName: &userName,
		Email:    &email,
		Password: &password,
	}
	headers := map[string]string{}
	w = util.PostJson(urlIndex, body, router, headers)
	assert.Equal(200, w.Code)
	resBody := w.Body.String()
	var baseRes vojo.BaseRes

	err := json.Unmarshal([]byte(resBody), &baseRes)
	assert.Equal(nil, err)
	assert.Equal(constants.NORMAL_RESPONSE_STATUS, baseRes.Rescode)
}
func TestMain(m *testing.M) {
	fmt.Println("test begin")
	InitRouter()
	m.Run()
	fmt.Println("test end")
}

func InitRouter() {
	r := gin.New()
	r.POST("/user", RegisterUser)
	router = r
}
