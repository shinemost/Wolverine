package models

import (
	"time"
)

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
	Id           int64     `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
	UpdateTime   time.Time `json:"update_time" db:"update_time"`
}

type Post struct {
	Status      int32     `json:"status" db:"status"`
	CommunityId int64     `json:"community_id" db:"community_id" binding:"required"`
	Id          int64     `json:"id,string" db:"post_id"`
	AuthorId    int64     `json:"author_id,string" db:"author_id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
	UpdateTime  time.Time `json:"update_time" db:"update_time"`
}

type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	*Community `json:"_community"`
	*Post      `json:"_post"`
}

type ParamLikeData struct {
	PostId    int64 `json:"post_id,string" binding:"required"`
	Direction int64 `json:"direction,string" binding:"required,oneof=1 -1"`
}

const (
	DirectionLike   = 1
	DirectionUnLike = -1
)
