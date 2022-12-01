package controllers

import (
	"hjfu/Wolverine/logic"
	"hjfu/Wolverine/models"
	"net/http"
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
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
		}
		return
	}

	err := logic.Register(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "注册成功")
}

func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("LoginHandler with invalid param", zap.Error(err))

		// 因为有的错误 比如json格式不对的错误 是不属于validator错误的 自然无法翻译，所以这里要做类型判断
		errs, ok := err.(validator.ValidationErrors)

		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)),
			})
		}
		return
	}

	err := logic.Login(p)
	if err != nil {
		// 可以在日志中 看出 到底是哪些用户一直在尝试登录
		zap.L().Error("login failed", zap.String("username", p.UserName), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码不正确",
		})
		return
	}
	c.JSON(http.StatusOK, "登录成功")
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
