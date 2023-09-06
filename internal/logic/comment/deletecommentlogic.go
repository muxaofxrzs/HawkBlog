package comment

import (
	"context"
	"hawk/internal/dao/mongo"

	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentReq) (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line
	err = mongo.DeleteComment(req)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "删除评论失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCode{
		Code:    types.OK,
		Message: "删除评论成功",
		Data:    struct{}{},
	}, nil
}
