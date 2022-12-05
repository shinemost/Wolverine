package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"hjfu/Wolverine/logic"
	"strconv"
)

func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("GetCommunityList error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	communityIDStr := c.Param("id")
	communityId, err := strconv.ParseInt(communityIDStr, 10, 64)
	if err != nil {
		zap.L().Error("GetCommunityListDetail error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityById(communityId)
	if err != nil {
		zap.L().Error("GetCommunityListDetail error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
