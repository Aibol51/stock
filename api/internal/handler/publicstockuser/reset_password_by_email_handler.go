package publicstockuser

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/publicstockuser"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /stockuser/reset_password_by_email publicstockuser ResetPasswordByEmail
//
// Reset password by Email | 通过邮箱重置密码
//
// Reset password by Email | 通过邮箱重置密码
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StockUserResetPasswordByEmailReq
//
// Responses:
//  200: BaseMsgResp

func ResetPasswordByEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StockUserResetPasswordByEmailReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := publicstockuser.NewResetPasswordByEmailLogic(r.Context(), svcCtx)
		resp, err := l.ResetPasswordByEmail(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
