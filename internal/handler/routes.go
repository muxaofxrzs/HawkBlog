// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	article "hawk/internal/handler/article"
	comment "hawk/internal/handler/comment"
	user "hawk/internal/handler/user"
	"hawk/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/Captcha",
				Handler: user.CaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/Register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/GetUserInformation",
				Handler: user.GetUserInformationHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/UpdateUserInformation",
				Handler: user.UpdateUserInformationHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/addArticle",
				Handler: article.AddArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/addDraft",
				Handler: article.AddDraftHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateArticle",
				Handler: article.UpdateArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/listArticle",
				Handler: article.ListArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/examineArticle",
				Handler: article.ExamineArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/examineArticles",
				Handler: article.ExamineArticlesHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/deleteArticle",
				Handler: article.DeleteArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/likeArticle",
				Handler: article.LikeArticleHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/article"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/article",
				Handler: comment.CreateCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/article",
				Handler: comment.DeleteCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/article",
				Handler: comment.UpdateCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/article",
				Handler: comment.GetCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/article/like",
				Handler: comment.PostCommentLikeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/articletoc",
				Handler: comment.CommentToCommentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/comment"),
	)
}
