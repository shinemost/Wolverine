package models

// 定义请求的参数结构体
type ParamRegister struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binging:"required"`
}

type User struct {
	UserId   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type Community struct {
	Id           int64  `json:"id" db:"community_id"`
	Name         string `json:"name" db:"community_name"`
	Introduction string `json:"introduction" db:"introduction"`
}
