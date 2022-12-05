package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"hjfu/Wolverine/logic"
	"hjfu/Wolverine/models"
	"strconv"
)

func CreatePostHandler(c *gin.Context) {
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CreatePostHandler error", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	authorId, err := getCurrentUserId(c)
	if err != nil {
		ResponseError(c, CodeNoLogin)
		return
	}
	p.AuthorId = authorId
	msg, err := logic.CreatePost(p)
	zap.L().Info("CreatePostHandlerSuccess", zap.String("postId", msg))
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, msg)
}

func PostDetailHandler(c *gin.Context) {
	postIDStr := c.Param("id")
	postId, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		zap.L().Error("GetPostDetail error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostById(postId)
	if err != nil {
		zap.L().Error("GetPostDetail error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetPostListHandler(c *gin.Context) {
	pageSizeStr := c.Query("pageSize")
	pageNumStr := c.Query("pageNum")

	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	pageNum, err := strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	if pageNum < 1 {
		ResponseErrorWithMsg(c, CodeInvalidParam, "页码不能小于1")
		return
	}
	data, err := logic.GetPostList(pageSize, pageNum)
	if err != nil {
		zap.L().Error("GetPostDetailHandler", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func PostLikeHandler(c *gin.Context) {
	p := new(models.ParamLikeData)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("PostLikeHandler with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	id, err := getCurrentUserId(c)
	if err != nil {
		ResponseError(c, CodeNoLogin)
		return
	}
	// 业务处理
	err = logic.PostLike(p, id)
	if err != nil {
		// 可以把
		if errors.Is(err, logic.ErrAleadyLike) || errors.Is(err, logic.ErrAleadyUnLike) {
			ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	ResponseSuccess(c, "success")
}

func GetPostListHandler2(c *gin.Context) {
	// 获取参数和参数校验
	p := new(models.ParamListData)
	// 校验下参数
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("CreatePostHandler with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
		} else {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	apiList, err := logic.GetPostList2(p)
	if err != nil {
		return
	}
	ResponseSuccess(c, apiList)
}
