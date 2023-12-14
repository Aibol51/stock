package Like

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/like"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLikeLogic {
	return &DeleteLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLikeLogic) DeleteLike(in *core.UUIDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.Like.Delete().Where(like.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).Exec(l.ctx)

	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
