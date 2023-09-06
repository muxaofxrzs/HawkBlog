package user

import (
	"context"
	"fmt"
	"hawk/internal/dao/mysql"
	"hawk/internal/tools"
	"strconv"

	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line
	dbRet, code := mysql.GetUser(req.Name, req.Password)
	fmt.Println(dbRet)
	if code == 1 {
		{
			return &types.HttpCode{
				Code:    types.UserInfoErr,
				Message: "登录失败",
			}, nil
		}
	}
	idInt64, _ := strconv.ParseInt(dbRet.Id, 10, 64)
	a, _ := tools.Token.GetToken(idInt64, dbRet.UserName, dbRet.Name, "user")
	if err != nil {
		resp = &types.HttpCode{
			Code:    types.UserInfoErr,
			Message: "生成Token失败",
			Data:    struct{}{},
		}
		return resp, nil
	}
	return &types.HttpCode{
		Code:    types.OK,
		Message: "登录成功",
		Data: Token{
			AccessToken:  a,
			RefreshToken: "",
		},
	}, nil
}
