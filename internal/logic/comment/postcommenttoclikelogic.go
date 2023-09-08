package comment

import (
	"context"
	"fmt"
	"hawk/internal/dao/mongo"
	"hawk/internal/dao/mysql"
	"hawk/internal/dao/redis"
	"strconv"

	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostCommenttocLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostCommenttocLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostCommenttocLikeLogic {
	return &PostCommenttocLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostCommenttocLikeLogic) PostCommenttocLike(req *types.PostCommenttocLikeReq) (resp *types.HttpCode, err error) {
	value := l.ctx.Value("userId")
	userId, _ := strconv.ParseInt(fmt.Sprintf("%s", value), 10, 64)
	//判断点赞的评论是否被删除了
	judge := mongo.CheckCommentExist(req.FirstCommentId, req.CommentId)
	if !judge {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "评论不存在",
			Data:    struct{}{},
		}, err
	}
	//判断用户点赞记录是否存在
	status, err := mysql.CheckLikeExist(userId, req.FirstCommentId, req.CommentId)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "评论信息点赞失败",
			Data:    struct{}{},
		}, err
	}
	//添加点赞记录
	err = mongo.PostCommenttocLike(status, req)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "评论信息点赞失败",
			Data:    struct{}{},
		}, err
	}
	err = mysql.PostCommentLike(status, userId, req.CommentId)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "评论信息点赞失败",
			Data:    struct{}{},
			//设置mongo中点赞的变化
		}, err
	}
	//在redis修改评论的热度
	err = redis.PostCommenttocLike(status, req)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "评论信息点赞失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCode{
		Code:    types.OK,
		Message: "评论信息点赞成功",
		Data:    struct{}{},
	}, nil
}
