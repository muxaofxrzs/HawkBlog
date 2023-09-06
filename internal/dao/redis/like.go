package redis

import (
	"HawkBlog/internal/types"
	"HawkBlog/model"
	"context"
	"fmt"
	"strconv"
)

func PostCommentLike(status int64, req *types.PostCommentLikeReq) (err error) {
	fmt.Println(req.CommentId)
	commentId := strconv.FormatInt(req.CommentId, 10)
	fmt.Println(commentId)
	key := model.ArticleCommentScore + strconv.FormatInt(req.ArticleId, 10)
	if status == 0 {
		_, err = ClientRe.ZIncrBy(context.Background(), key, 432.0, commentId).Result()
		return
	}
	_, err = ClientRe.ZIncrBy(context.Background(), key, -432.0, commentId).Result()
	return
}
