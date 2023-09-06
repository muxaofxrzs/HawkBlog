package mongo

import (
	"HawkBlog/internal/types"
	"HawkBlog/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strconv"
	"time"
)

// 删除二级评论
func DeleteCommenttoc(req *types.DeleteCommenttocReq) (err error) {
	collection := ClientMo.Database("hawk").Collection(strconv.FormatInt(req.FirstCommentId, 10))
	filter := bson.M{"commentid": req.CommentId}
	update := bson.D{
		{"$set", bson.D{
			{"status", 1},
			// 添加其他字段的更新
		}},
	}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	return err
}

// 修改二级评论
func UpdateCommenttoc(req *types.UpdateCommenttocReq) (err error) {
	collection := ClientMo.Database("hawk").Collection(strconv.FormatInt(req.FirstCommentId, 10))
	filter := bson.M{"commentid": req.CommentId}
	update := bson.D{
		{"$set", bson.D{
			{"comment", req.Comment},
			{"updatetime", time.Now()},
			// 添加其他字段的更新
		}},
	}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	return err
}

// 获取一级评论的所有二级评论
func GetAllCommenttoc(firstCommentId int64, commentIdList []int64) (data []model.ArticleComment, err error) {
	collection := ClientMo.Database("hawk").Collection(strconv.FormatInt(firstCommentId, 10))
	// 构建查询条件
	for _, v := range commentIdList {
		filter := bson.M{"status": 0, "commentid": v}
		var result model.ArticleComment
		err = collection.FindOne(context.Background(), filter).Decode(&result)
		if err != nil {
			log.Fatalf("查询错误：%v", err)
		}
		data = append(data, result)
	}
	return
}
