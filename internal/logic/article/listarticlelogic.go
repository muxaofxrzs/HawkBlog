package article

import (
	"context"
	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListArticleLogic {
	return &ListArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListArticleLogic) ListArticle() (resp *types.HttpCodeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
