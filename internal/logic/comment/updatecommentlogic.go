package comment

import (
	"HawkBlog/internal/dao/mongo"
	"context"

	"HawkBlog/internal/svc"
	"HawkBlog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCommentLogic) UpdateComment(req *types.UpdateCommentReq) (resp *types.HttpCode, err error) {
	// todo: add your logic here and delete this line
	err = mongo.UpdateComment(req)
	if err != nil {
		return &types.HttpCode{
			Code:    types.DoErr,
			Message: "修改评论失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCode{
		Code:    types.OK,
		Message: "修改评论成功",
		Data:    struct{}{},
	}, nil
}
