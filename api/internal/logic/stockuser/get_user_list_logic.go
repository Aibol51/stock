package stockuser

import (
	"context"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetUserListLogic) GetUserList(req *types.StockUserListReq) (resp *types.StockUserListResp, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.CoreRpc.GetStockUserList(l.ctx, &core.StockUserListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Username: req.Username,
		Nickname: req.Nickname,
		Email:    req.Email,
		Mobile:   req.Mobile,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.StockUserListResp{}
	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data, types.StockUserInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			Username:    v.Username,
			Nickname:    v.Nickname,
			Mobile:      v.Mobile,
			Email:       v.Email,
			Avatar:      v.Avatar,
			Status:      v.Status,
			Description: v.Description,
			HomePath:    v.HomePath,
		})
	}
	resp.Data.Total = data.Total
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	return resp, nil
}
