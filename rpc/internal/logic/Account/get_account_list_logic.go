package Account

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/account"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountListLogic {
	return &GetAccountListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountListLogic) GetAccountList(in *core.AccountListReq) (*core.AccountListResp, error) {
	var predicates []predicate.Account
	if in.AccountName != nil {
		predicates = append(predicates, account.AccountNameContains(*in.AccountName))
	}
	if in.Password != nil {
		predicates = append(predicates, account.PasswordContains(*in.Password))
	}
	if in.Email != nil {
		predicates = append(predicates, account.EmailContains(*in.Email))
	}
	result, err := l.svcCtx.DB.Account.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.AccountListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.AccountInfo{
			Id:          pointy.GetPointer(v.ID.String()),
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:      pointy.GetPointer(uint32(v.Status)),
			AccountName: &v.AccountName,
			Password:    &v.Password,
			Mobile:      v.Mobile,
			Email:       v.Email,
			Avatar:      &v.Avatar,
			Gender:      pointy.GetPointer(string(v.Gender)),
		})
	}

	return resp, nil
}
