package comment

import (
	"context"
	"fmt"
	"hawk/internal/dao/mongo"
	"hawk/internal/pkg/snowflake"
	"hawk/model"
	"strconv"
	"time"

	"hawk/internal/svc"
	"hawk/internal/types"

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
		CommentId:  snowflake.GenId(),
		Comment:    req.Comment,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Like:       0,
		Status:     0,
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
	return &types.HttpCode{
		Code:    types.OK,
		Message: "添加评论信息成功",
		Data:    comment,
	}, nil
	return
}
