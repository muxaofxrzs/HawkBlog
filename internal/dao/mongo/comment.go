package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"hawk/internal/types"
	"hawk/model"
	"log"
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
func GetAllComment(articleId int64, commentIdList []int64) (data []model.ArticleComment, err error) {
	collection := ClientMo.Database("hawk").Collection(strconv.FormatInt(articleId, 10))
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
func CheckCommentExist(articleId, commentId int64) (judge bool) {
	collection := ClientMo.Database("hawk").Collection(strconv.FormatInt(articleId, 10))
	filter := bson.M{"commentid": commentId, "status": 0}
	count, err := collection.CountDocuments(context.Background(), filter)
	fmt.Println(count)
	if err != nil {
		fmt.Println("err:", err)
	}
	if count != 1 {
		return false
	}
	return true
}
func GetCommentCount(req *types.GetCommentCountReq) (count int64, err error) {
	collection := ClientMo.Database("hawk").Collection(strconv.FormatInt(req.RequireId, 10))
	filter := bson.M{"status": 0}
	count, err = collection.CountDocuments(context.Background(), filter)
	return count, err
}
