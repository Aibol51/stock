package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type View struct {
	ent.Schema
}

func (View) Fields() []ent.Field {
	return []ent.Field{}
}

func (View) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

func (View) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("views").
			Unique(),
		edge.From("post", Post.Type).
			Ref("views"),
	}
}
