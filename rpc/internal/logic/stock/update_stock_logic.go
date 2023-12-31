package stock

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStockLogic {
	return &UpdateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStockLogic) UpdateStock(in *core.StockInfo) (*core.BaseResp, error) {
	query := l.svcCtx.DB.Stock.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
		SetNotNilStockName(in.StockName).
		SetNotNilStockCode(in.StockCode).
		SetNotNilIsRecommend(in.IsRecommend)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}
	// 处理新增的上涨和下跌字段
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

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
