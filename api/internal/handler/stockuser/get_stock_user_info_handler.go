package stockuser

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/stockuser"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
)

// swagger:route get /stockuser/info stockuser GetStockUserInfo
//
// Get stockuser basic information | 获取用户基本信息
//
// Get stockuser basic information | 获取用户基本信息
//
// Responses:
//  200: StockUserBaseIDInfoResp

func GetStockUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := stockuser.NewGetStockUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetStockUserInfo()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
