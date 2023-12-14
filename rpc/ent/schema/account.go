package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
	mixins2 "github.com/suyuan32/simple-admin-core/rpc/ent/schema/mixins"
)

type Account struct {
	ent.Schema
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_name").Unique().Comment("登录名").Annotations(entsql.WithComments(true)),
		field.String("password").Comment("密码").Annotations(entsql.WithComments(true)),
		field.Uint32("mobile").Nillable().Comment("手机号").Annotations(entsql.WithComments(true)),
		field.String("email").Nillable().Comment("邮箱").Annotations(entsql.WithComments(true)),
		field.String("avatar").
			SchemaType(map[string]string{dialect.MySQL: "varchar(512)"}).
			Optional().
			Default("").
			Comment("Avatar | 头像路径").
			Annotations(entsql.WithComments(true)),
		field.Enum("gender").Values("UNKOWN", "MALE", "FEMALE").Default("UNKOWN").Comment("性别").Annotations(entsql.WithComments(true)),
	}
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("comments", Comment.Type),
		edge.To("likes", Like.Type),
		edge.To("views", View.Type),
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_name").
			Unique(),
	}
}
