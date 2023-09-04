package user

import (
	"context"

	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInformationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInformationLogic {
	return &UpdateUserInformationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInformationLogic) UpdateUserInformation(req *types.UpdateUserInformationReq) (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line

	return
}
