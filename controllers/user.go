package controllers

import (
	"errors"
	"hjfu/Wolverine/dao/mysql"
	"hjfu/Wolverine/logic"
	"hjfu/Wolverine/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func RegisterHandler(c *gin.Context) {
	p := new(models.ParamRegister)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("RegisterHandler with invalid param", zap.Error(err))

		// 因为有的错误 比如json格式不对的错误 是不属于validator错误的 自然无法翻译，所以这里要做类型判断
		errs, ok := err.(validator.ValidationErrors)

		if !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	// 业务处理
	err := logic.Register(p)
	if err != nil {
		zap.L().Error("register failed", zap.String("username", p.UserName), zap.Error(err))
		if errors.Is(err, mysql.UserAleadyExists) {
			ResponseError(c, CodeUserExist)
		} else {
			ResponseError(c, CodeInvalidParam)
		}
		return
	}
	ResponseSuccess(c, "注册成功")
}

func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("LoginHandler with invalid param", zap.Error(err))

		// 因为有的错误 比如json格式不对的错误 是不属于validator错误的 自然无法翻译，所以这里要做类型判断
		errs, ok := err.(validator.ValidationErrors)

		if !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	token, err := logic.Login(p)
	if err != nil {
		// 可以在日志中 看出 到底是哪些用户一直在尝试登录
		zap.L().Error("login failed", zap.String("username", p.UserName), zap.Error(err))
		if errors.Is(err, mysql.WrongPassword) {
			ResponseError(c, CodeInvalidPassword)
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	ResponseSuccess(c, token)
}

// 去除报错信息中的结构体信息
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		// 这里 算法非常简单 就是遍历你的错误信息 然后把key值取出来 把.之前的信息去掉就行了
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func PingHandler(c *gin.Context) {
	// 这里post man 模拟的 将token auth-token
	zap.L().Debug("ping", zap.String("ping-username", c.GetString("username")))
	ResponseSuccess(c, "pong")

}

func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("GetCommunityList error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
