package logic

import (
	"go.uber.org/zap"
	"hjfu/Wolverine/dao/mysql"
	"hjfu/Wolverine/dao/redis"
	"hjfu/Wolverine/models"
	"hjfu/Wolverine/pkg/snowflake"
	"strconv"
)

func CreatePost(p *models.Post) (msg string, err error) {
	p.Id = snowflake.GenId()
	zap.L().Debug("createPostLogic", zap.Int64("postId", p.Id))
	err = mysql.InsertPost(p)
	if err != nil {
		return "failed", err
	}
	// 去点赞数量的 zset 新增一条记录
	err = redis.AddPost(p.Id)
	if err != nil {
		return "", err
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

	return rangeInitApiPostDetail(postList)

}

func GetPostList2(params *models.ParamListData) (apiPostDetailList []*models.ApiPostDetail, err error) {
	// 最热
	if params.Order == models.OrderByHot {
		// 先去redis 里面取 最新的数据
		ids, err := redis.GetPostIdsByScore(params.PageSize, params.PageNum)
		if err != nil {
			return nil, err
		}
		postLists, err := mysql.GetPostListByIds(ids)
		if err != nil {
			return nil, err
		}
		return rangeInitApiPostDetail(postLists)

	} else if params.Order == models.OrderByTime {
		//最新
		return GetPostList(params.PageSize, params.PageNum)
	}
	return nil, nil
}

func rangeInitApiPostDetail(posts []*models.Post) (apiPostDetailList []*models.ApiPostDetail, err error) {
	for _, post := range posts {
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
