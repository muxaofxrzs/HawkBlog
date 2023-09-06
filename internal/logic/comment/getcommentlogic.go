package comment

import (
	"context"
	"hawk/internal/dao/mongo"

	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentLogic) GetComment(req *types.GetAllCommentReq) (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line
	data, err := mongo.GetAllComment(req)
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
