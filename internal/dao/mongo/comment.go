package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"hawk/internal/types"
	"hawk/model"
	"strconv"
	"time"
)

// 添加评论
func CreateComment(articleId int64, req *model.ArticleComment) error {
	database := ClientMo.Database("hawk")
	collection := database.Collection(strconv.FormatInt(articleId, 10))
	_, err := collection.InsertOne(context.Background(), req)
	if err != nil {
		fmt.Println("评论数据添加失败")
	}
	return err
}

// 获取指定博客的所有评论
func GetAllComment(req *types.GetAllCommentReq) (data []model.ArticleComment, err error) {
	collection := ClientMo.Database("hawk").Collection(strconv.FormatInt(req.ArticleId, 10))
	// 构建查询条件
	filter := bson.M{"status": 0}
	// 执行查询操作
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("mongo数据查询失败")
		return data, err
	}
	// 遍历查询结果
	for cur.Next(context.Background()) {
		var d model.ArticleComment
		err = cur.Decode(&d)
		data = append(data, d)
	}
	return
}

// 删除指定博客的指定评论
func DeleteComment(req *types.DeleteCommentReq) (err error) {
	collection := ClientMo.Database("hawk").Collection(strconv.FormatInt(req.ArticleId, 10))
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

// 修改指定博客的指定评论
func UpdateComment(req *types.UpdateCommentReq) (err error) {
	collection := ClientMo.Database("hawk").Collection(strconv.FormatInt(req.ArticleId, 10))
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
func CommentToComment(CommentId int64, req *model.ArticleComment) error {
	database := ClientMo.Database("hawk")
	collection := database.Collection(strconv.FormatInt(CommentId, 10))
	_, err := collection.InsertOne(context.Background(), req)
	if err != nil {
		fmt.Println("评论数据添加失败")
	}
	return err
}
