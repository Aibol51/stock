package stockuser

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStockUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockUserInfoLogic {
	return &GetStockUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetStockUserInfoLogic) GetStockUserInfo() (resp *types.StockUserBaseIDInfoResp, err error) {
	// todo: add your logic here and delete this line
	stockUser, err := l.svcCtx.CoreRpc.GetStockUserById(l.ctx, &core.UUIDReq{Id: l.ctx.Value("userId").(string)})
	if err != nil {
		return nil, err
	}
	return &types.StockUserBaseIDInfoResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
		Data: types.StockUserBaseIDInfo{
			UUID:        stockUser.Id,
			Username:    stockUser.Username,
			Nickname:    stockUser.Nickname,
			Avatar:      stockUser.Avatar,
			HomePath:    stockUser.HomePath,
			Description: stockUser.Description,
		},
	}, nil
}
