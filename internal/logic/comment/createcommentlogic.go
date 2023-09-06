package comment

import (
	"context"
	"fmt"
	"hawk/internal/dao/mongo"
	"hawk/internal/dao/redis"
	"hawk/internal/pkg/snowflake"
	"hawk/model"
	"strconv"
	"time"

	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentReq) (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line
	value := l.ctx.Value("userId")
	userId, _ := strconv.ParseInt(fmt.Sprintf("%s", value), 10, 64)

	//var userId int64
	//userId = l.ctx.Value("id").(int64)
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
	err = mongo.CreateComment(req.ArticleId, comment)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "添加评论信息失败",
			Data:    struct{}{},
		}, err
	}
	//将评论ID存放在redis中
	err = redis.CreateComment(req.ArticleId, comment)
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
