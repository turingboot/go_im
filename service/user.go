package service

import (
	"fmt"
	"go_im/common/global"
	"go_im/dal/model"
	"go_im/dal/query"
	util "go_im/utils"
	"math/rand"
	"time"
)

type UserService struct{}

// UserRegister 用户注册
func (s *UserService) UserRegister(mobile, plainPwd, nickname, avatar, sex string) (model.User, error) {
	registerUser := model.User{}
	u := query.Use(global.Db).User
	qUser, err := u.Where(u.Mobile.Eq(mobile)).First()

	//如果用户已经注册,返回错误信息
	if err != nil && qUser != nil {
		return model.User{}, err
	}

	registerUser.Mobile = mobile
	registerUser.Avatar = avatar
	registerUser.Nickname = nickname
	registerUser.Sex = sex
	registerUser.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	registerUser.Passwd = util.MakePasswd(plainPwd, registerUser.Salt)
	registerUser.Createat = time.Now()

	//插入用户信息
	err = u.Create(&registerUser)
	return registerUser, err
}

// Login 用户登录
func (s *UserService) Login(mobile, plainPwd string) (model.User, error) {
	//数据库操作
	loginUser := model.User{}
	u := query.Use(global.Db).User
	qUser, err := u.Where(u.Mobile.Eq(mobile)).First()
	loginUser = *qUser
	//用户不存在
	if qUser == nil {
		return loginUser, err
	}
	//判断密码是否正确
	if !util.ValidatePasswd(plainPwd, loginUser.Salt, loginUser.Passwd) {
		return loginUser, err
	}
	//刷新用户登录的token值
	token := util.GenRandomStr(32)
	loginUser.Token = token

	_, err = u.Where(u.ID.Eq(loginUser.ID)).Update(u.Token, token)
	if err != nil {
		return loginUser, err
	}

	//返回新用户信息
	return loginUser, nil
}

// Find 查找某个用户
func (s *UserService) Find(userId int64) *model.User {
	u := query.Use(global.Db).User
	qUser, _ := u.Where(u.ID.Eq(userId)).First()
	return qUser
}
