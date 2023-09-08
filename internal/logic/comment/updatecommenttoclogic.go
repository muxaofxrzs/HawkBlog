package comment

import (
	"context"
	"hawk/internal/dao/mongo"

	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommenttocLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCommenttocLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommenttocLogic {
	return &UpdateCommenttocLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCommenttocLogic) UpdateCommenttoc(req *types.UpdateCommenttocReq) (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line
	err = mongo.UpdateCommenttoc(req)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "修改评论失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCode{
		Code:    types.OK,
		Message: "修改评论成功",
		Data:    struct{}{},
	}, nil
}
