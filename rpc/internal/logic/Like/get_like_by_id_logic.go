package Like

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLikeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLikeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLikeByIdLogic {
	return &GetLikeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLikeByIdLogic) GetLikeById(in *core.UUIDReq) (*core.LikeInfo, error) {
	result, err := l.svcCtx.DB.Like.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.LikeInfo{
		Id:        pointy.GetPointer(result.ID.String()),
		CreatedAt: pointy.GetPointer(result.CreatedAt.Unix()),
		UpdatedAt: pointy.GetPointer(result.UpdatedAt.Unix()),
		Status:    pointy.GetPointer(uint32(result.Status)),
	}, nil
}
