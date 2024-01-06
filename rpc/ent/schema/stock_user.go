package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"

	// "entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"

	mixins2 "github.com/suyuan32/simple-admin-core/rpc/ent/schema/mixins"
)

// StockUser holds the schema definition for the StockUser entity.
type StockUser struct {
	ent.Schema
}

// Fields of the StockUser.
func (StockUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Optional().Unique().
			Comment("StockUser's login name | 登录名").
			Annotations(entsql.WithComments(true)),
		field.String("password").
			Comment("Password | 密码").
			Annotations(entsql.WithComments(true)),
		field.String("nickname").Unique().
			Comment("Nickname | 昵称").
			Annotations(entsql.WithComments(true)),
		field.String("description").Optional().
			Comment("The description of stock user | 股票用户的描述信息").
			Annotations(entsql.WithComments(true)),
		field.String("home_path").Default("/dashboard").
			Comment("The home page that the stock user enters after logging in | 股票用户登陆后进入的首页").
			Annotations(entsql.WithComments(true)),
		field.String("mobile").Optional().Unique().
			Comment("Mobile number | 手机号").
			Annotations(entsql.WithComments(true)),
		field.String("email").Optional().
			Comment("Email | 邮箱号").
			Annotations(entsql.WithComments(true)),
		field.Text("avatar").
			Optional().
			Default("").
			Comment("Avatar | 头像路径").
			Annotations(entsql.WithComments(true)),

		// Add additional fields specific to StockUser here
		field.String("last_login_info").
			Optional().
			Default("").
			Comment("Last login information | 最后登录信息").
			Annotations(entsql.WithComments(true)),
	}
}

// Mixin of the StockUser.
func (StockUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

// Indexes of the StockUser.
func (StockUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "email", "mobile").
			Unique(),
		// Add additional indexes specific to StockUser here

	}
}

// Annotations of the StockUser.
func (StockUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sys_stock_users"},
	}
}
