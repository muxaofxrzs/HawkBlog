package user

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"hawk/internal/dao/mysql"
	"hawk/internal/svc"
	"hawk/internal/types"
	"reflect"
	"strconv"
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

	value := l.ctx.Value("userId")
	fmt.Println(reflect.TypeOf(value))
	userId, _ := strconv.ParseInt(fmt.Sprintf("%s", value), 10, 64)
	ret := mysql.GetUserInfo(userId)
	fmt.Println("++++++++++++++++++++++++")
	if ret.Id == 0 {
		return &types.HttpCode{
			Code:    types.OK,
			Message: "登录成功",
			Data:    struct{}{},
		}, nil
	}
	return &types.HttpCode{
		Code:    types.OK,
		Message: "登录成功",
		Data:    ret,
	}, nil

}
