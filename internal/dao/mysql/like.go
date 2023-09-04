package mysql

import (
	"context"
	"fmt"
	"hawk/internal/types"
)

func CheckLikeExist(userId int64, req *types.PostCommentLikeReq) (status int64, err error) {
	sqlStr := "select status from comment_like where user_id = ? and comment_id = ?"
	err = GlobalConn.QueryRowCtx(context.Background(), &status, sqlStr, userId, req.CommentId)
	if err != nil {
		fmt.Println("执行了创建语句")
		sqlStr1 := "insert into comment_like(user_id,article_id,comment_id,status) values (?,?,?,?)"
		_, err = GlobalConn.ExecCtx(context.Background(), sqlStr1, userId, req.ArticleId, req.CommentId, 1)
	}
	return status, err
}

func PostCommentLike(status, userId int64, req *types.PostCommentLikeReq) (err error) {
	//用户未点赞就添加一条点赞记录
	if status == 0 {
		sqlStr := "update comment_like set status = 1 where user_id = ? and comment_id = ?"
		_, err = GlobalConn.ExecCtx(context.Background(), sqlStr, userId, req.CommentId)
		return err
	}
	//用户点过赞的话就将用户设置status为0
	sqlStr := "update comment_like set status = 0 where user_id = ? and comment_id = ?"
	_, err = GlobalConn.ExecCtx(context.Background(), sqlStr, userId, req.CommentId)
	return err
}
