package article

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"hawk/internal/dao/mysql"
	"hawk/internal/svc"
	"hawk/internal/types"
)

type ExamineArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExamineArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExamineArticlesLogic {
	return &ExamineArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 主页面获取热度最高的六篇文章
func (l *ExamineArticlesLogic) ExamineArticles() (resp *types.HttpCodeResp, err error) {
	// todo: add your logic here and delete this line
	//获取前六篇的文章
	articles, err := mysql.GetTopSixArticles()
	return &types.HttpCodeResp{
		Code:    1000,
		Message: "获取前六章文章成功",
		Data:    articles,
	}, err
}
