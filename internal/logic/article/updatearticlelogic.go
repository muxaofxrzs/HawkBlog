package article

import (
	"context"
	"fmt"

	"github.com/russross/blackfriday/v2"
	"hawk/internal/dao/mysql"
	"hawk/internal/svc"
	"hawk/internal/types"
	"hawk/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleLogic) UpdateArticle(req *types.UpdateArticleReq) (resp *types.HttpCodeResp, err error) {
	// todo: add your logic here and delete this line
	contentstr := req.Content
	content := string(blackfriday.Run([]byte(contentstr)))
	articleId := req.ArticleId
	article := model.Article{
		ArticleId: articleId,
		Content:   content,
	}
	err = mysql.UpdataEssay(article)
	fmt.Println("=====成功")
	if err != nil {
		return &types.HttpCodeResp{
			Code:    1008,
			Message: "失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCodeResp{
		Code:    1009,
		Message: "成功",
		Data: struct {
		}{},
	}, nil
}
