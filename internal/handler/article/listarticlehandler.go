package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hawk/internal/logic/article"
	"hawk/internal/svc"
)

func ListArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := article.NewListArticleLogic(r.Context(), svcCtx)
		resp, err := l.ListArticle()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
