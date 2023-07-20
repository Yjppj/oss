package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sso/controller/jwt"
	"sso/model"
	"sso/response"
)

type ParamLogin struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	//1.获取请求参数及参数校验
	var p ParamLogin
	if err := c.ShouldBindJSON(&p); err != nil { //这个地方只能判断是不是json格式的数据
		response.FailWithMessage("参数有误，登录失败", c)
		return
	}
	//2.业务逻辑处理
	//3.返回响应
	user, err := (&model.DingUser{Mobile: p.Mobile, Password: p.Password}).Login()
	if err != nil {
		response.FailWithMessage("登录失败",c)
		return
	}
	// 生成JWT
	token, err := jwt.GenToken(c, user)
	if err != nil {
		zap.L().Debug("JWT生成错误")
		return
	}
	user.AuthToken = token

	if err != nil {
		response.FailWithMessage("用户登录失败", c)
	} else {
		response.OkWithDetailed(user, "用户登录成功", c)
	}
}
