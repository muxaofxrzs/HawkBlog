package article

import (
	"context"
	"fmt"
	"hawk/internal/dao/mysql"
	"hawk/internal/dao/redis"
	"strconv"

	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeArticleLogic {
	return &LikeArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeArticleLogic) LikeArticle(req *types.LikeReq) (resp *types.HttpCodeResp, err error) {
	// todo: add your logic here and delete this line
	value := l.ctx.Value("userId")
	userId, _ := strconv.ParseInt(fmt.Sprintf("%s", value), 10, 64)
	articleId := req.ArticleId
	//在redis中查询该用户是否点赞过
	liked, err := redis.CheckUserLike(articleId, userId)
	if err != nil {
		return &types.HttpCodeResp{
			Code:    1000,
			Message: "查询点赞记录失败",
			Data:    struct{}{},
		}, err
	}
	if liked {
		return &types.HttpCodeResp{
			Code:    1000,
			Message: "已经点赞过",
			Data:    struct{}{},
		}, err
	}
	likeCount, err := mysql.LikeArticleInMySQL(articleId)
	if err != nil {
		return &types.HttpCodeResp{
			Code:    1000,
			Message: "点赞失败",
			Data:    struct{}{},
		}, err
	}
	//将这次的点赞记录记录到redis中
	err = redis.StoreUserLike(articleId, userId)
	if err != nil {
		return &types.HttpCodeResp{
			Code:    1000,
			Message: "点赞存入redis失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCodeResp{
		Code:    1000,
		Message: "点赞成功",
		Data:    likeCount,
	}, err
}
