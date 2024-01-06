package publicstockuser

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/publicstockuser"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /stockuser/login_by_sms publicstockuser LoginBySms
//
// Log in by SMS | 短信登录
//
// Log in by SMS | 短信登录
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StockUserLoginBySmsReq
//
// Responses:
//  200: StockUserLoginResp

func LoginBySmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StockUserLoginBySmsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publicstockuser.NewLoginBySmsLogic(r.Context(), svcCtx)
		resp, err := l.LoginBySms(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
