package Post

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostByIdLogic {
	return &GetPostByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostByIdLogic) GetPostById(in *core.UUIDReq) (*core.PostInfo, error) {
	result, err := l.svcCtx.DB.Post.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.PostInfo{
		Id:        pointy.GetPointer(result.ID.String()),
		CreatedAt: pointy.GetPointer(result.CreatedAt.Unix()),
		UpdatedAt: pointy.GetPointer(result.UpdatedAt.Unix()),
		Status:    pointy.GetPointer(uint32(result.Status)),
		Title:     &result.Title,
		Content:   &result.Content,
		Author:    &result.Author,
		Category:  &result.Category,
		Tags:      &result.Tags,
		Summary:   &result.Summary,
		Cover:     &result.Cover,
	}, nil
}
