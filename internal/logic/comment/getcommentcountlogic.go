package comment

import (
	"HawkBlog/internal/dao/mongo"
	"context"

	"HawkBlog/internal/svc"
	"HawkBlog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentCountLogic {
	return &GetCommentCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentCountLogic) GetCommentCount(req *types.GetCommentCountReq) (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line
	count, err := mongo.GetCommentCount(req)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "获取数量失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCode{
		Code:    types.OK,
		Message: "获取数量成功",
		Data:    count,
	}, nil
}
