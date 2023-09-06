package comment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hawk/internal/logic/comment"
	"hawk/internal/svc"
	"hawk/internal/types"
)

func PostCommenttocLikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostCommenttocLikeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewPostCommenttocLikeLogic(r.Context(), svcCtx)
		resp, err := l.PostCommenttocLike(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
