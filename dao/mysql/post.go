package mysql

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"hjfu/Wolverine/models"
	"strconv"
	"strings"
)

func InsertPost(p *models.Post) error {
	sqlStr := "insert into post(post_id,title,content,author_id,community_id) values (?,?,?,?,?)"
	_, err := db.Exec(sqlStr, p.Id, p.Title, p.Content, p.AuthorId, p.CommunityId)
	if err != nil {
		zap.L().Error("create post error", zap.Error(err))
		return err
	}
	return nil
}

func GetPostById(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := "select post_id,title,content,author_id,community_id," +
		"status,create_time from post where post_id=?"
	err = db.Get(post, sqlStr, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("no post")
			err = nil
		}
	}
	return post, err
}

func GetPostList(offset int64, pageSize int64) (posts []*models.Post, err error) {
	zap.L().Info("GetPostList", zap.String("offset", strconv.FormatInt(offset, 10)), zap.String("pageSize", strconv.FormatInt(pageSize, 10)))
	sqlStr := "select post_id,title,content,author_id,community_id,create_time,update_time " +
		" from post order by create_time desc limit ?,?"
	posts = make([]*models.Post, 0, pageSize)
	err = db.Select(&posts, sqlStr, offset, pageSize)
	if err != nil {
		return nil, err
	}
	return posts, err

}

func GetPostListByIds(ids []string) (postList []*models.Post, err error) {
	//FIND_IN_SET 按照给定的顺序 来返回结果集
	sqlStr := "select post_id,title,content,author_id,community_id,create_time,update_time" +
		" from post where post_id in (?) order by FIND_IN_SET(post_id,?)"
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	if err != nil {
		zap.L().Error("GetPostListByIds", zap.Error(err))
		return nil, err
	}
	return postList, nil
}
