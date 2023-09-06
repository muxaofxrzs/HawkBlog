package article

import (
	"context"
	"hawk/internal/dao/mysql"

	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(req *types.DeleteReq) (resp *types.HttpCodeResp, err error) {
	// todo: add your logic here and delete this line
	articleId := req.ArticleId

	err = mysql.DeleteEssay(articleId)
	if err != nil {
		return &types.HttpCodeResp{
			Code:    1002,
			Message: "删除文章失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCodeResp{
		Code:    1003,
		Message: "删除文章成功",
		Data:    struct{}{},
	}, err
}
