package article

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"hawk/internal/dao/mysql"
	"hawk/internal/pkg/snowflake"
	"hawk/internal/svc"
	"hawk/internal/types"
	"hawk/model"
	"strconv"
)

type AddArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddArticleLogic {
	return &AddArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddArticleLogic) AddArticle(req *types.AddArticleReq) (resp *types.HttpCodeResp, err error) {
	// todo: add your logic here and delete this line
	title := req.Title
	content := req.Content
	label := req.Label
	value := l.ctx.Value("userId")
	userId, _ := strconv.ParseInt(fmt.Sprintf("%s", value), 10, 64)
	article := model.Article{
		ArticleId: snowflake.GenId(),
		UserId:    userId,
		Title:     title,
		Content:   content,
		Label:     label,
	}
	err = mysql.AddEssay(article)
	fmt.Println("=====成功")
	if err != nil {
		return &types.HttpCodeResp{
			Code:    1000,
			Message: "失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCodeResp{
		Code:    1001,
		Message: "成功",
		Data:    article,
	}, nil

}
