package logic

import (
	"hjfu/Wolverine/dao/mysql"
	"hjfu/Wolverine/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	return mysql.GetCommunityList()
}

func GetCommunityById(CommunityId int64) (model *models.Community, err error) {
	return mysql.GetCommunityById(CommunityId)
}
