package comment

import (
	"context"
	"hawk/internal/dao/mongo"
	"hawk/internal/dao/redis"

	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommenttocLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommenttocLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommenttocLogic {
	return &DeleteCommenttocLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommenttocLogic) DeleteCommenttoc(req *types.DeleteCommenttocReq) (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line
	err = mongo.DeleteCommenttoc(req)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "删除评论失败",
			Data:    struct{}{},
		}, err
	}
	err = redis.DeleteCommenttoc(req)
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
