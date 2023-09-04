package article

import (
	"context"
	"fmt"
	"hawk/internal/dao/mysql"
	"hawk/internal/svc"
	"hawk/internal/types"

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
	content := req.Content
	articleId := req.ArticleId
	err = mysql.UpdataEssay(articleId, content)
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
