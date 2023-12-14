package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
	mixins2 "github.com/suyuan32/simple-admin-core/rpc/ent/schema/mixins"
)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique().Comment("标题"),
		field.String("content").Comment("内容"),
		field.String("author").Comment("作者"),
		field.String("category").Comment("分类"),
		field.String("tags").Comment("标签"),
		field.String("summary").Comment("摘要"),
		field.String("cover").Comment("封面"),
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).Ref("posts").Unique(),
		edge.To("comments", Comment.Type),
		edge.To("likes", Like.Type),
		edge.To("views", View.Type),
	}
}

func (Post) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title").
			Unique(),
	}
}
