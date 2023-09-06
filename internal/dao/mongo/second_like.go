package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"hawk/internal/types"
	"strconv"
)

func PostCommenttocLike(status int64, req *types.PostCommenttocLikeReq) (err error) {
	if status == 0 {
		articleId := strconv.FormatInt(req.FirstCommentId, 10)
		collection := ClientMo.Database("hawk").Collection(articleId)
		filterInc := bson.M{"commentid": req.CommentId}
		updateInc := bson.M{"$inc": bson.M{"like": 1}}
		_, err = collection.UpdateOne(context.Background(), filterInc, updateInc)
		return
	}
	articleId := strconv.FormatInt(req.FirstCommentId, 10)
	collection := ClientMo.Database("hawk").Collection(articleId)
	filterInc := bson.M{"commentid": req.CommentId}
	updateInc := bson.M{"$inc": bson.M{"like": -1}}
	_, err = collection.UpdateOne(context.Background(), filterInc, updateInc)
	return
}
