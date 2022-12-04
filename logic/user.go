package logic

import (
	"hjfu/Wolverine/dao/mysql"
	"hjfu/Wolverine/models"
	"hjfu/Wolverine/pkg/jwt"
	"hjfu/Wolverine/pkg/snowflake"
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
