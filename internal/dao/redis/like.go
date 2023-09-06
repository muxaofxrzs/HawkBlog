package redis

import (
<<<<<<< HEAD
	"hawk/internal/types"
	"hawk/model"
	"context"
	"fmt"
=======
	"context"
	"fmt"
	"hawk/internal/types"
	"hawk/model"
>>>>>>> 2cb677a71339b35127a74f83095f18be7858e549
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
