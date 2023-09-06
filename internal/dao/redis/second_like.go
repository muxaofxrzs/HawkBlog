package redis

import (
	"context"
	"fmt"
	"hawk/internal/types"
	"hawk/model"
	"strconv"
)

func PostCommenttocLike(status int64, req *types.PostCommenttocLikeReq) (err error) {
	commentId := strconv.FormatInt(req.CommentId, 10)
	fmt.Println(commentId)
	key := model.CommentCommentScore + strconv.FormatInt(req.FirstCommentId, 10)
	if status == 0 {
		_, err = ClientRe.ZIncrBy(context.Background(), key, 432.0, commentId).Result()
		return
	}
	_, err = ClientRe.ZIncrBy(context.Background(), key, -432.0, commentId).Result()
	return
}
