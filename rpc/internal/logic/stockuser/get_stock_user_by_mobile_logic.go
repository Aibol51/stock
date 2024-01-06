package stockuser

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-core/rpc/ent/stockuser" // Import the correct package for StockUser

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStockUserByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStockUserByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStockUserByMobileLogic {
	return &GetStockUserByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStockUserByMobileLogic) GetStockUserByMobile(in *core.StockUserMobileReq) (*core.StockUserInfo, error) {
	// todo: add your logic here and delete this line
	result, err := l.svcCtx.DB.StockUser.Query().Where(stockuser.MobileEQ(in.Mobile)).First(l.ctx)
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
