package article

import (
	"context"
	"fmt"
	"hawk/internal/dao/mysql"
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

func (l *ListArticleLogic) ListArticle(req *types.ListTitleReq) (resp *types.HttpCodeResp, err error) {
	// todo: add your logic here and delete this line
	//用户在输入文章题目的时候，暂时根据模糊查询，后续使用es进行查询，返回响应的文章的题目
	title := req.Title
	articleId := req.ArticleId
	fmt.Println(articleId)
	//首先判断是不是第一页，如果传进来的articleid是0的话，说明是第一页，首先获取第一个的id
	if articleId == 0 {
		articleId = mysql.GetTopOneTitle(title)

	} else {
		articleId = req.ArticleId
	}
	articles, err := mysql.GetNextTitleId(articleId, req.PageSize, title)
	if err != nil {
		return &types.HttpCodeResp{
			Code:    1010,
			Message: "查看文章题目失败",
			Data: struct {
			}{},
		}, nil
	}
	return &types.HttpCodeResp{
		Code:    1009,
		Message: "查看文章题目成功",
		Data:    articles,
	}, nil
}
