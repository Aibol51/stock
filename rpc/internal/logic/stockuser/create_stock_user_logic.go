package stockuser

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStockUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStockUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStockUserLogic {
	return &CreateStockUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStockUserLogic) CreateStockUser(in *core.StockUserInfo) (*core.BaseUUIDResp, error) {
	// Assuming the structure of StockUserInfo is similar to UserInfo
	result, err := l.svcCtx.DB.StockUser.Create().
		SetNotNilUsername(in.Username).
		SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(*in.Password))).
		SetNotNilNickname(in.Nickname).
		SetNotNilEmail(in.Email).
		SetNotNilMobile(in.Mobile).
		SetNotNilAvatar(in.Avatar).
		SetNotNilHomePath(in.HomePath).
		SetNotNilDescription(in.Description).
		SetNotNilLastLoginInfo(in.LastLoginInfo).
		Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
