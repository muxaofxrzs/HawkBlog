package article

import (
	"context"
	"hawk/internal/dao/mysql"
	"hawk/model"

	"hawk/internal/svc"
	"hawk/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExamineArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExamineArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExamineArticleLogic {
	return &ExamineArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExamineArticleLogic) ExamineArticle(req *types.ExamineReq) (resp *types.HttpCodeResp, err error) {
	// todo: add your logic here and delete this line
	articleId := req.ArticleId
	var article model.Article
	err, article = mysql.FindEssay(articleId)
	if err != nil {
		return &types.HttpCodeResp{
			Code:    1004,
			Message: "查询失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCodeResp{
		Code:    1004,
		Message: "查询成功",
		Data:    article,
	}, nil
}
