package mysql

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
	"hjfu/Wolverine/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id,community_name,introduction from community"
	err = db.Select(&communityList, sqlStr)
	if err != nil {
		// 空数据的时候 不算错误 只是没有板块而已
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("no community ")
			err = nil
		}
	}
	return
}
func GetCommunityById(id int64) (community *models.Community, err error) {
	community = new(models.Community)
	sqlStr := "select community_id,community_name,introduction,create_time,update_time " +
		"from community where community_id=?"
	err = db.Get(community, sqlStr, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("no community")
			err = nil
		}
	}
	return community, err
}
