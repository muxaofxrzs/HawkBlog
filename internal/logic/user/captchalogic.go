package user

import (
	"context"
	"fmt"
	"hawk/internal/dao/redis"
	"hawk/internal/tools"
	"math/rand"
	"time"

	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) Captcha(req *types.CaptchaReq) (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf(req.Email)
	code := generateCode()
	fmt.Println(code)
	err = tools.SendEmail(req.Email, code)
	if err != nil {
		return &types.HttpCode{
			Code:    types.NotFound,
			Message: "发送验证码失败,尝试重新发送",
		}, err
	}
	errRedis := redis.ClientRe.Set(context.Background(), req.Email, code, 0).Err()
	if errRedis != nil {
		// 处理存储出错的情况
	}
	return &types.HttpCode{
		Code:    types.OK,
		Message: "发送成功",
		Data:    struct{}{},
	}, nil

}

func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(999999)
	return fmt.Sprintf("%06d", code)
}
