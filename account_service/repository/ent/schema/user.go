package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Positive().Unique(),
		field.String("nick_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(20)", // Override MySQL.
		}).NotEmpty().Unique().Comment("名称"),

		field.String("password").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).NotEmpty().Comment("密码").Sensitive(),

		field.String("avatar").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).NotEmpty().Comment("头像"),

		field.Int8("status").SchemaType(map[string]string{
			dialect.MySQL: "tinyint unsigned", // Override MySQL.
		}).NonNegative().Default(0).Comment("状态,是否被封，0：正常，大于0被封"),

		field.Int64("score").SchemaType(map[string]string{
			dialect.MySQL: "int UNSIGNED", // Override MySQL.
		}).NonNegative().Default(0).Comment("分数"),

		field.Int8("role").SchemaType(map[string]string{
			dialect.MySQL: "tinyint UNSIGNED", // Override MySQL.
		}).NonNegative().Default(0).Comment("用户的角色，0：普通用户，1：admin"),

		field.Int64("create_at").SchemaType(map[string]string{
			dialect.MySQL: "int UNSIGNED", // Override MySQL.
		}).NonNegative().Default(0).Comment("创建时间").Immutable(),

		field.Int64("update_at").SchemaType(map[string]string{
			dialect.MySQL: "int UNSIGNED", // Override MySQL.
		}).NonNegative().Default(0).Comment("最后更新时间"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}
