package home_advertise

import (
	"net/http"

	"github.com/feihua/zero-admin/api/admin/internal/logic/sms/home_advertise"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryHomeAdvertiseDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryHomeAdvertiseDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := home_advertise.NewQueryHomeAdvertiseDetailLogic(r.Context(), svcCtx)
		resp, err := l.QueryHomeAdvertiseDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
