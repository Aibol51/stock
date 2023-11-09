package stock

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/ent/stock"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockListLogic {
	return &GetStockListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockListLogic) GetStockList(in *core.StockListReq) (*core.StockListResp, error) {
	var predicates []predicate.Stock
	if in.StockName != nil {
		predicates = append(predicates, stock.StockNameContains(*in.StockName))
	}
	if in.StockCode != nil {
		predicates = append(predicates, stock.StockCodeContains(*in.StockCode))
	}
	result, err := l.svcCtx.DB.Stock.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.StockListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.StockInfo{
			Id:          pointy.GetPointer(v.ID.String()),
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:      pointy.GetPointer(uint32(v.Status)),
			StockName:   &v.StockName,
			StockCode:   &v.StockCode,
			IsRecommend: &v.IsRecommend,
		})
	}

	return resp, nil
}
