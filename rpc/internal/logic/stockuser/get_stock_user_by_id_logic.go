package stockuser

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	// "github.com/suyuan32/simple-admin-core/rpc/ent"
	"github.com/suyuan32/simple-admin-core/rpc/ent/stockuser" // Import the correct package for StockUser

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockUserByIdLogic {
	return &GetStockUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockUserByIdLogic) GetStockUserById(in *core.UUIDReq) (*core.StockUserInfo, error) {
	result, err := l.svcCtx.DB.StockUser.Query().Where(stockuser.IDEQ(uuidx.ParseUUIDString(in.Id))).First(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.StockUserInfo{
		Id:            pointy.GetPointer(result.ID.String()),
		Username:      &result.Username,
		Password:      &result.Password, // Note: Usually, we don't retrieve the password
		Nickname:      &result.Nickname,
		Description:   &result.Description,
		HomePath:      &result.HomePath,
		Mobile:        &result.Mobile,
		Email:         &result.Email,
		Avatar:        &result.Avatar,
		CreatedAt:     pointy.GetPointer(result.CreatedAt.Unix()),
		UpdatedAt:     pointy.GetPointer(result.UpdatedAt.Unix()),
		Status:        pointy.GetPointer(uint32(result.Status)),
		LastLoginInfo: pointy.GetPointer(result.LastLoginInfo),
		// Add other fields as necessary
	}, nil
}

// Remove or modify the helper functions like GetRoleIds, GetRoleNames, etc., as they are not used here.
