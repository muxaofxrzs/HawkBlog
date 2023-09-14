package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"hawk/internal/types"
	"hawk/model"
	"strconv"
	"time"
)

func CreateCommenttoc(firstCommentId int64, req *model.ArticleComment) (err error) {
	timeKey := model.CommentCommentTime + strconv.FormatInt(firstCommentId, 10)
	scoreKey := model.CommentCommentScore + strconv.FormatInt(firstCommentId, 10)
	pipeline := ClientRe.TxPipeline()
	pipeline.ZAdd(context.Background(), timeKey, &redis.Z{
		Member: req.CommentId,
		Score:  float64(time.Now().Unix()),
	})
	pipeline.ZAdd(context.Background(), scoreKey, &redis.Z{
		Member: req.CommentId,
		Score:  0,
	})
	_, err = pipeline.Exec(context.Background())
	return err
}
func DeleteCommenttoc(req *types.DeleteCommenttocReq) (err error) {
	_, err = ClientRe.ZRem(context.Background(), model.CommentCommentScore+strconv.FormatInt(req.FirstCommentId, 10), req.CommentId).Result()
	if err != nil {
		fmt.Println("redis删除出错：", err)
	}
	_, err = ClientRe.ZRem(context.Background(), model.CommentCommentTime+strconv.FormatInt(req.FirstCommentId, 10), req.CommentId).Result()
	if err != nil {
		fmt.Println("redis删除出错：", err)
	}
	return err
}
func GetCommenttoc(req *types.GetCommenttocReq) (commentIdList []int64, err error) {
	commentKey := model.CommentCommentTime + strconv.FormatInt(req.FirstCommentId, 10)
	if req.Method == "score" {
		commentKey = model.CommentCommentScore + strconv.FormatInt(req.FirstCommentId, 10)
	}
	var rank int64 = -1
	if req.LastCommentId != 0 {
		rank, err = ClientRe.ZRevRank(context.Background(), commentKey, strconv.FormatInt(req.LastCommentId, 10)).Result()
	}
	start := rank + 1
	stop := rank + req.PageNumber
	result, err := ClientRe.ZRevRangeWithScores(context.Background(), commentKey, start, stop).Result()
	for _, v := range result {
		data, _ := strconv.ParseInt(v.Member.(string), 10, 64)
		commentIdList = append(commentIdList, data)
	}
	return commentIdList, err
}
