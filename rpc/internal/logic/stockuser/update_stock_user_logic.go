package stockuser

import (
	"context"

	"github.com/suyuan32/simple-admin-common/enum/common"

	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/suyuan32/simple-admin-core/rpc/internal/logic/token"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/entx"

	"github.com/suyuan32/simple-admin-core/rpc/ent"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type UpdateStockUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStockUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStockUserLogic {
	return &UpdateStockUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStockUserLogic) UpdateStockUser(in *core.StockUserInfo) (*core.BaseResp, error) {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		updateQuery := tx.StockUser.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
			SetNotNilUsername(in.Username).
			SetNotNilNickname(in.Nickname).
			SetNotNilEmail(in.Email).
			SetNotNilMobile(in.Mobile).
			SetNotNilAvatar(in.Avatar).
			SetNotNilHomePath(in.HomePath).
			SetNotNilDescription(in.Description).
			SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
			SetNotNilLastLoginInfo(in.LastLoginInfo)

		if in.Password != nil {
			updateQuery = updateQuery.SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(*in.Password)))
		}

		// Removed logic for roles and positions as they are not part of StockUser

		if in.Password != nil || (in.Status != nil && uint8(*in.Status) != common.StatusNormal) {
			_, err := token.NewBlockUserAllTokenLogic(l.ctx, l.svcCtx).BlockUserAllToken(&core.UUIDReq{Id: *in.Id})
			if err != nil {
				return err
			}
		}

		return updateQuery.Exec(l.ctx)
	})
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{
		Msg: i18n.Success,
	}, nil
}

// Note: Adjust this implementation according to the specific needs of your StockUser entity.
