package stockuser

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/ent/stockuser" // Import the correct package for StockUser

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockUserListLogic {
	return &GetStockUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockUserListLogic) GetStockUserList(in *core.StockUserListReq) (*core.StockUserListResp, error) {
	var predicates []predicate.StockUser

	if in.Mobile != nil {
		predicates = append(predicates, stockuser.MobileEQ(*in.Mobile))
	}

	if in.Username != nil {
		predicates = append(predicates, stockuser.UsernameContains(*in.Username))
	}

	if in.Email != nil {
		predicates = append(predicates, stockuser.EmailEQ(*in.Email))
	}

	if in.Nickname != nil {
		predicates = append(predicates, stockuser.NicknameContains(*in.Nickname))
	}

	// Add other predicates as necessary, based on the fields of StockUser

	stockUsers, err := l.svcCtx.DB.StockUser.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.StockUserListResp{}
	resp.Total = stockUsers.PageDetails.Total

	for _, v := range stockUsers.List {
		resp.Data = append(resp.Data, &core.StockUserInfo{
			Id:            pointy.GetPointer(v.ID.String()),
			Avatar:        &v.Avatar,
			Mobile:        &v.Mobile,
			Email:         &v.Email,
			Status:        pointy.GetPointer(uint32(v.Status)),
			Username:      &v.Username,
			Nickname:      &v.Nickname,
			HomePath:      &v.HomePath,
			Description:   &v.Description,
			CreatedAt:     pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:     pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			LastLoginInfo: pointy.GetPointer(v.LastLoginInfo),
			// Add other fields as necessary
		})
	}

	return resp, nil
}

// Remove helper functions related to roles and positions if they are not used here.
