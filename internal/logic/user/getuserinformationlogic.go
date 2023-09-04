package user

import (
	"context"
	"fmt"
	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInformationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInformationLogic {
	return &GetUserInformationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInformationLogic) GetUserInformation() (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line
	// 从上下文中获取用户信息
	fmt.Println("______________________________")
	value := l.ctx.Value("userName")
	fmt.Print(value)
	return
}
