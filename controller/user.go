package controller

import (
	"go_im/dal/model"
	"go_im/service"
	"go_im/utils"
	"net/http"
)

var UserService service.UserService

// UserRegister 用户注册
func UserRegister(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	err := util.Bind(request, &user)
	if err != nil {
		return
	}
	user, err = UserService.UserRegister(user.Mobile, user.Passwd, user.Nickname, user.Avatar, user.Sex)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}
}

// UserLogin 用户登录
func UserLogin(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		return
	}

	mobile := request.PostForm.Get("mobile")
	plainPwd := request.PostForm.Get("passwd")

	//校验参数
	if len(mobile) == 0 || len(plainPwd) == 0 {
		util.RespFail(writer, "用户名或密码不正确")
	}

	loginUser, err := UserService.Login(mobile, plainPwd)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, loginUser, "")
	}
}
