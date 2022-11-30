package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hjfu/Wolverine/logic"
	"hjfu/Wolverine/models"
	"net/http"
)

func RegisterHandler(c *gin.Context) {
	p := new(models.ParamRegister)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("RegisterHandler with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	if len(p.UserName) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名不能为空",
		})
		return
	}
	if len(p.Password) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "密码不能为空",
		})
		return
	}
	if p.Password != p.RePassword {
		c.JSON(http.StatusOK, gin.H{
			"msg": "两次密码输入不一致",
		})
		return
	}
	logic.Register(p)
	c.JSON(http.StatusOK, "注册成功")
}
