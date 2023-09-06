package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"hawk/model"
	"strconv"
	"time"
)

func CreateComment(articleId int64, req *model.ArticleComment) (err error) {
	timeKey := model.ArticleCommentTime + strconv.FormatInt(articleId, 10)
	scoreKey := model.ArticleCommentScore + strconv.FormatInt(articleId, 10)
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
