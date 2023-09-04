package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hawk/internal/logic/article"
	"hawk/internal/svc"
	"hawk/internal/types"
)

func ExamineArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExamineReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewExamineArticleLogic(r.Context(), svcCtx)
		resp, err := l.ExamineArticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
