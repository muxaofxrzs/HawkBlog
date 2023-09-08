package article

import (
	"context"
	"fmt"

	"github.com/russross/blackfriday/v2"

	"github.com/zeromicro/go-zero/core/logx"
	"hawk/internal/dao/mysql"
	"hawk/internal/pkg/snowflake"
	"hawk/internal/svc"
	"hawk/internal/types"
	"hawk/model"
	"strconv"

	"time"
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
	//输入的文章内容是一个markdown的类型的
	contentstr := req.Content
	content := string(blackfriday.Run([]byte(contentstr)))
	label := req.Label
	value := l.ctx.Value("userId")
	userId, _ := strconv.ParseInt(fmt.Sprintf("%s", value), 10, 64)
	userNameAny := l.ctx.Value("userName")
	userName := fmt.Sprintf("%s", userNameAny)
	article := model.Article{
		ArticleId:  snowflake.GenId(),
		UserId:     userId,
		Title:      title,
		Content:    content,
		Label:      label,
		UserName:   userName,
		StartTime:  time.Now(),
		UpdataTime: time.Now(),
		Status:     1,
		Heat:       0,
	}
	err = mysql.AddEssay(article)
	fmt.Println("=====成功")
	if err != nil {
		return &types.HttpCodeResp{
			Code:    1000,
			Message: "发布文章失败",
			Data:    struct{}{},
		}, err
	}
	return &types.HttpCodeResp{
		Code:    1001,
		Message: "发布文章成功",
		Data:    article,
	}, nil

}
