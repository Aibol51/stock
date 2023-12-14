package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Like struct {
	ent.Schema
}

func (Like) Fields() []ent.Field {
	return []ent.Field{}
}

func (Like) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("likes").
			Unique(),
		edge.From("post", Post.Type).
			Ref("likes"),
	}
}
