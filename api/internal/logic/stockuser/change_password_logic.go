package stockuser

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.StockUserChangePasswordReq) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line
	stockUserData, err := l.svcCtx.CoreRpc.GetStockUserById(l.ctx, &core.UUIDReq{Id: l.ctx.Value("userId").(string)})
	if err != nil {
		return nil, err
	}
	if encrypt.BcryptCheck(req.OldPassword, *stockUserData.Password) {
		result, err := l.svcCtx.CoreRpc.UpdateStockUser(l.ctx, &core.StockUserInfo{
			Id:       pointy.GetPointer(l.ctx.Value("userId").(string)),
			Password: pointy.GetPointer(req.NewPassword),
		})
		if err != nil {
			return nil, err
		}

		return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
	}
	return nil, errorx.NewCodeInvalidArgumentError("login.wrongPassword")
}
