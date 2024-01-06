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
		// 添加上涨字段
		field.String("stock_rise").
			Optional().
			Comment("Stock rise value").
			Annotations(entsql.WithComments(true)),
		// 添加下跌字段
		field.String("stock_fall").
			Optional().
			Comment("Stock fall value").
			Annotations(entsql.WithComments(true)),
		// 添加股票添加时间字段
		field.String("add_time").
			Optional().
			Comment("Stock addition time as a string").
			Annotations(entsql.WithComments(true)),
		// 添加股票详情信息字段
		field.String("details").
			Optional().
			Comment("Detailed information about the stock").
			Annotations(entsql.WithComments(true)),
		// 添加股票标签字段
		field.String("stock_tags").
			Optional().
			Comment("Tags associated with the stock").
			Annotations(entsql.WithComments(true)),
	}
}

func (Stock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}
