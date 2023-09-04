package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hawk/internal/logic/user"
	"hawk/internal/svc"
)

func GetUserInformationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserInformationLogic(r.Context(), svcCtx)
		resp, err := l.GetUserInformation()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
