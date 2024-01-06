package stock

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStockLogic {
	return &CreateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStockLogic) CreateStock(in *core.StockInfo) (*core.BaseUUIDResp, error) {
	query := l.svcCtx.DB.Stock.Create().
		SetNotNilStockName(in.StockName).
		SetNotNilStockCode(in.StockCode).
		SetNotNilIsRecommend(in.IsRecommend)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}
	// Handle the new stock rise and fall fields
	if in.StockRise != nil {
		query.SetNotNilStockRise(in.StockRise)
	}
	if in.StockFall != nil {
		query.SetNotNilStockFall(in.StockFall)
	}
	if in.AddTime != nil {
		query.SetNotNilAddTime(in.AddTime)
	}
	if in.Details != nil {
		query.SetNotNilDetails(in.Details)
	}
	if in.StockTags != nil {
		query.SetNotNilStockTags(in.StockTags)
	}

	result, err := query.Save(l.ctx)

	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
