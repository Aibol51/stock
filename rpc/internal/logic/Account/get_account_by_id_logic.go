package Account

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountByIdLogic {
	return &GetAccountByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountByIdLogic) GetAccountById(in *core.UUIDReq) (*core.AccountInfo, error) {
	result, err := l.svcCtx.DB.Account.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.AccountInfo{
		Id:          pointy.GetPointer(result.ID.String()),
		CreatedAt:   pointy.GetPointer(result.CreatedAt.Unix()),
		UpdatedAt:   pointy.GetPointer(result.UpdatedAt.Unix()),
		Status:      pointy.GetPointer(uint32(result.Status)),
		AccountName: &result.AccountName,
		Password:    &result.Password,
		Mobile:      result.Mobile,
		Email:       result.Email,
		Avatar:      &result.Avatar,
		Gender:      pointy.GetPointer(string(result.Gender)),
	}, nil
}
