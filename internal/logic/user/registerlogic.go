package user

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"hawk/internal/dao/mysql"
	"hawk/internal/dao/redis"

	"hawk/internal/svc"
	"hawk/internal/types"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line
	fmt.Print("_______________________________________")
	//client := tools.CreateRedisClient()
	code, err := redis.ClientRe.Get(context.Background(), req.Email).Result()
	fmt.Println("code", code)
	if err == nil {
		// 键不存在的处理逻辑
	} else if err != nil {
		// 获取代码出错的处理逻辑
	} else {
		// 使用获取到的代码进行后续操作
	}
	if code != req.Code {
		return &types.HttpCode{
			Code:    types.UserInfoErr,
			Message: "验证码不正确",
			Data:    struct{}{},
		}, nil
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++")
	ret := mysql.Registered(req)
	if ret == 1 {
		return &types.HttpCode{
			Code:    types.OK,
			Message: "用户已存在",
			Data:    struct{}{},
		}, nil
	}
	if ret == 2 {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "注册失败,请重新注册",
			Data:    struct{}{},
		}, nil
	}
	return &types.HttpCode{
		Code:    types.OK,
		Message: "注册成功",
		Data:    struct{}{},
	}, nil
}
