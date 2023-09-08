package comment

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hawk/internal/dao/mongo"
	"hawk/internal/dao/redis"
	"hawk/internal/svc"
	"hawk/internal/types"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentReq) (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line
	//在mongo删除指定评论信息
	err = mongo.DeleteComment(req)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "删除评论失败",
			Data:    struct{}{},
		}, err
	}
	err = redis.DeleteComment(req)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "删除评论失败",
			Data:    struct{}{},
		}, err
	}

	return &types.HttpCode{
		Code:    types.OK,
		Message: "删除评论成功",
		Data:    struct{}{},
	}, nil
}
