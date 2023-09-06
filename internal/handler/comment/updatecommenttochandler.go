package comment

import (
	"net/http"

	"HawkBlog/internal/logic/comment"
	"HawkBlog/internal/svc"
	"HawkBlog/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateCommenttocHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCommenttocReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewUpdateCommenttocLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCommenttoc(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
