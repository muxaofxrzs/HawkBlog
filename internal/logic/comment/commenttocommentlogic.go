package comment

import (
	"HawkBlog/internal/dao/mongo"
	"HawkBlog/internal/dao/redis"
	"HawkBlog/internal/pkg/snowflake"
	"HawkBlog/model"
	"context"
	"fmt"
	"strconv"
	"time"

	"HawkBlog/internal/svc"
	"HawkBlog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentToCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentToCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentToCommentLogic {
	return &CommentToCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentToCommentLogic) CommentToComment(req *types.CommentToCommentReq) (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line
	value := l.ctx.Value("userId")
	userId, _ := strconv.ParseInt(fmt.Sprintf("%s", value), 10, 64)
	comment := &model.ArticleComment{
		UserId:     userId,
		CommentId:  snowflake.GenID(),
		Comment:    req.Comment,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Like:       0,
		Status:     0,
	}
	//判断一级评论是否存在
	judge := mongo.CheckCommentExist(req.ArticleId, req.CommmentId)
	if !judge {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "评论不存在",
			Data:    struct{}{},
		}, err
	}
	//将评论信息存储在mongo中
	err = mongo.CommentToComment(req.CommmentId, comment)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "添加评论信息失败",
			Data:    struct{}{},
		}, err
	}
	err = redis.CreateCommenttoc(req.CommmentId, comment)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "添加评论信息失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCode{
		Code:    types.OK,
		Message: "添加评论信息成功",
		Data:    comment,
	}, nil
}
