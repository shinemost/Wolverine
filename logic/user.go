package logic

import (
	"errors"
	"go.uber.org/zap"
	"hjfu/Wolverine/dao/mysql"
	"hjfu/Wolverine/dao/redis"
	"hjfu/Wolverine/models"
	"hjfu/Wolverine/pkg/jwt"
	"hjfu/Wolverine/pkg/snowflake"
	"strconv"
)

var ErrAleadyLike = errors.New("不能重复点赞")
var ErrAleadyUnLike = errors.New("不能重复点踩")

func Register(register *models.ParamRegister) (err error) {

	// 判断用户是否存在
	err = mysql.CheckUserExist(register.UserName)
	if err != nil {
		// db 出错
		return err
	}
	// 生成userid
	userId := snowflake.GenId()
	// 构造一个User db对象
	user := models.User{
		UserId:   userId,
		Username: register.UserName,
		Password: register.Password,
	}
	// 保存数据库
	err = mysql.InsertUser(&user)
	if err != nil {
		return err
	}
	return

}

func Login(login *models.ParamLogin) (
	string,
	error,
) {
	user := models.User{
		Username: login.UserName,
		Password: login.Password,
	}
	if err := mysql.Login(&user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.Username, user.UserId)
}
func GetCommunityList() (communityList []*models.Community, err error) {
	return mysql.GetCommunityList()
}

func GetCommunityById(CommunityId int64) (model *models.Community, err error) {
	return mysql.GetCommunityById(CommunityId)
}

func CreatePost(p *models.Post) (msg string, err error) {
	p.Id = snowflake.GenId()
	zap.L().Debug("createPostLogic", zap.Int64("postId", p.Id))
	err = mysql.InsertPost(p)
	if err != nil {
		return "failed", err
	}
	return strconv.FormatInt(p.Id, 10), nil
}

func GetPostById(postId int64) (apiPostDetail *models.ApiPostDetail, err error) {

	post, err := mysql.GetPostById(postId)
	username, err := mysql.GetUserNameById(post.AuthorId)
	if err != nil {
		zap.L().Warn("no author ")
		err = nil
	}
	community, err := GetCommunityById(post.CommunityId)
	if err != nil {
		zap.L().Warn("no community ")
		err = nil
	}
	apiPostDetail = new(models.ApiPostDetail)
	apiPostDetail.AuthorName = username
	apiPostDetail.Post = post
	apiPostDetail.Community = community
	return apiPostDetail, err
}

func GetPostList(pageSize, pageNum int64) (apiPostDetailList []*models.ApiPostDetail, err error) {

	offset := pageSize * (pageNum - 1)
	postList, err := mysql.GetPostList(offset, pageSize)
	if err != nil {
		return nil, err
	}
	apiPostDetailList = make([]*models.ApiPostDetail, 0, 2)
	for _, post := range postList {
		//再查 作者 名称
		username, err := mysql.GetUserNameById(post.AuthorId)
		if err != nil {
			zap.L().Warn("no author ")
			err = nil
			return nil, err
		}
		//再查板块实体
		community, err := GetCommunityById(post.CommunityId)
		if err != nil {
			zap.L().Warn("no community ")
			err = nil
			return nil, err
		}
		apiPostDetail := new(models.ApiPostDetail)
		apiPostDetail.AuthorName = username
		apiPostDetail.Community = community
		apiPostDetail.Post = post
		apiPostDetailList = append(apiPostDetailList, apiPostDetail)
	}
	return apiPostDetailList, nil
}

func PostLike(postData *models.ParamLikeData, userId int64) error {
	// 查询之前有没有点过赞
	direction, flag := redis.CheckLike(postData.PostId, userId)

	if flag {
		// 如果之前点过赞 则要判断 这次是否是重复点赞
		if direction == postData.Direction && direction == models.DirectionLike {
			return ErrAleadyLike
		}
		// 如果之前点踩 则要判断 这次是否是重复 点踩
		if direction == postData.Direction && direction == models.DirectionUnLike {
			return ErrAleadyUnLike
		}
	}

	err := redis.DoLike(postData.PostId, userId, postData.Direction)
	if err != nil {
		return err
	}

	return nil
}
