package stock

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockByIdLogic {
	return &GetStockByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockByIdLogic) GetStockById(in *core.UUIDReq) (*core.StockInfo, error) {
	result, err := l.svcCtx.DB.Stock.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.StockInfo{
		Id:          pointy.GetPointer(result.ID.String()),
		CreatedAt:   pointy.GetPointer(result.CreatedAt.Unix()),
		UpdatedAt:   pointy.GetPointer(result.UpdatedAt.Unix()),
		Status:      pointy.GetPointer(uint32(result.Status)),
		StockName:   &result.StockName,
		StockCode:   &result.StockCode,
		IsRecommend: &result.IsRecommend,
	}, nil
}
