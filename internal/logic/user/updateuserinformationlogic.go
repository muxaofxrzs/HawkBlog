package user

import (
	"context"
	"fmt"
	"hawk/internal/dao/mysql"
	"reflect"
	"strconv"

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
	value := l.ctx.Value("userId")
	fmt.Println(reflect.TypeOf(value))
	userId, _ := strconv.ParseInt(fmt.Sprintf("%s", value), 10, 64)
	ret := mysql.UpdateUserInfo(req, userId)
	if ret == 1 {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "更新失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCode{
		Code:    types.OK,
		Message: "更新成功",
		Data:    struct{}{},
	}, nil
}
