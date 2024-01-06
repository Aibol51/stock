package publicstockuser

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/publicstockuser"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /stockuser/login_Mobile publicstockuser LoginByMobile
//
// Log in by mobile phone | 手机号登录
//
// Log in by mobile phone | 手机号登录
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StockUserLoginByMobileReq
//
// Responses:
//  200: StockUserLoginResp

func LoginByMobileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StockUserLoginByMobileReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publicstockuser.NewLoginByMobileLogic(r.Context(), svcCtx)
		resp, err := l.LoginByMobile(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
