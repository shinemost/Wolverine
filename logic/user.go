package logic

import (
	"go.uber.org/zap"
	"hjfu/Wolverine/dao/mysql"
	"hjfu/Wolverine/models"
	"hjfu/Wolverine/pkg/jwt"
	"hjfu/Wolverine/pkg/snowflake"
	"strconv"
)

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
