package stockuser

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateUserLogic) UpdateUser(req *types.StockUserInfo) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.CoreRpc.UpdateStockUser(l.ctx,
		&core.StockUserInfo{
			Id:          req.Id,
			Status:      req.Status,
			Username:    req.Username,
			Password:    req.Password,
			Nickname:    req.Nickname,
			Description: req.Description,
			HomePath:    req.HomePath,
			Mobile:      req.Mobile,
			Email:       req.Email,
			Avatar:      req.Avatar,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
