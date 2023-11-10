package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Stock struct {
	ent.Schema
}

func (Stock) Fields() []ent.Field {
	return []ent.Field{
		field.String("stock_name").
			Comment("Stock name").
			Annotations(entsql.WithComments(true)),
		field.String("stock_code").
			Comment("Stock code").
			Annotations(entsql.WithComments(true)),
		field.Bool("is_recommend").
			Default(false).
			Comment("Stock code").
			Annotations(entsql.WithComments(true)),
	}
}

func (Stock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}
