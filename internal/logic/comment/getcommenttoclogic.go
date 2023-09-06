package comment

import (
	"HawkBlog/internal/dao/mongo"
	"HawkBlog/internal/dao/redis"
	"HawkBlog/internal/svc"
	"HawkBlog/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommenttocLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommenttocLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommenttocLogic {
	return &GetCommenttocLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommenttocLogic) GetCommenttoc(req *types.GetCommenttocReq) (resp *types.HttpCode, err error) {
	commentIdList, err := redis.GetCommenttoc(req)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "获取评论信息失败",
			Data:    struct{}{},
		}, err
	}
	data, err := mongo.GetAllCommenttoc(req.FirstCommentId, commentIdList)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "获取评论信息失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCode{
		Code:    types.OK,
		Message: "获取评论信息成功",
		Data:    data,
	}, nil
}
