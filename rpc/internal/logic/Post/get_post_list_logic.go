package Post

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/post"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/errorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostListLogic {
	return &GetPostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostListLogic) GetPostList(in *core.PostListReq) (*core.PostListResp, error) {
	var predicates []predicate.Post
	if in.Title != nil {
		predicates = append(predicates, post.TitleContains(*in.Title))
	}
	if in.Content != nil {
		predicates = append(predicates, post.ContentContains(*in.Content))
	}
	if in.Author != nil {
		predicates = append(predicates, post.AuthorContains(*in.Author))
	}
	result, err := l.svcCtx.DB.Post.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.PostListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.PostInfo{
			Id:        pointy.GetPointer(v.ID.String()),
			CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:    pointy.GetPointer(uint32(v.Status)),
			Title:     &v.Title,
			Content:   &v.Content,
			Author:    &v.Author,
			Category:  &v.Category,
			Tags:      &v.Tags,
			Summary:   &v.Summary,
			Cover:     &v.Cover,
		})
	}

	return resp, nil
}
