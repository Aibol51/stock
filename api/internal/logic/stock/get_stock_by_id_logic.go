package stock

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStockByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockByIdLogic {
	return &GetStockByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStockByIdLogic) GetStockById(req *types.UUIDReq) (resp *types.StockInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetStockById(l.ctx, &core.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.StockInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.StockInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:      data.Status,
			StockName:   data.StockName,
			StockCode:   data.StockCode,
			IsRecommend: data.IsRecommend,
			StockRise:   data.StockRise,
			StockFall:   data.StockFall,
			AddTime:     data.AddTime,
			Details:     data.Details,
			StockTags:   data.StockTags,
		},
	}, nil
}
