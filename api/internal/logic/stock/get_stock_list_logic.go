package stock

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStockListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockListLogic {
	return &GetStockListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStockListLogic) GetStockList(req *types.StockListReq) (resp *types.StockListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetStockList(l.ctx,
		&core.StockListReq{
			Page:      req.Page,
			PageSize:  req.PageSize,
			StockName: req.StockName,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.StockListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.StockInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status:      v.Status,
				StockName:   v.StockName,
				StockCode:   v.StockCode,
				IsRecommend: v.IsRecommend,
			})
	}
	return resp, nil
}
