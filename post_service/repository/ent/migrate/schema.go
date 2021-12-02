// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommentColumns holds the columns for the "comment" table.
	CommentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "user_id", Type: field.TypeInt64, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "post_id", Type: field.TypeInt64, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "content", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(1024)"}},
		{Name: "status", Type: field.TypeInt8, Default: 0, SchemaType: map[string]string{"mysql": "tinyint unsigned"}},
		{Name: "floor", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "create_at", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "post_comments", Type: field.TypeInt64, Nullable: true, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
	}
	// CommentTable holds the schema information for the "comment" table.
	CommentTable = &schema.Table{
		Name:       "comment",
		Columns:    CommentColumns,
		PrimaryKey: []*schema.Column{CommentColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comment_post_comments",
				Columns:    []*schema.Column{CommentColumns[7]},
				RefColumns: []*schema.Column{PostColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "comment_user_id_post_id",
				Unique:  false,
				Columns: []*schema.Column{CommentColumns[1], CommentColumns[2]},
			},
		},
	}
	// PostColumns holds the columns for the "post" table.
	PostColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "user_id", Type: field.TypeInt64, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "topic", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(1024)"}},
		{Name: "status", Type: field.TypeInt8, Default: 0, SchemaType: map[string]string{"mysql": "tinyint unsigned"}},
		{Name: "reply_num", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "create_at", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "update_at", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
	}
	// PostTable holds the schema information for the "post" table.
	PostTable = &schema.Table{
		Name:       "post",
		Columns:    PostColumns,
		PrimaryKey: []*schema.Column{PostColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "post_user_id",
				Unique:  false,
				Columns: []*schema.Column{PostColumns[1]},
			},
		},
	}
	// ReplyColumns holds the columns for the "reply" table.
	ReplyColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "user_id", Type: field.TypeInt64, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "post_id", Type: field.TypeInt64, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "comment_id", Type: field.TypeInt64, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "parent_id", Type: field.TypeInt64, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "content", Type: field.TypeString, SchemaType: map[string]string{"mysql": "varchar(1024)"}},
		{Name: "status", Type: field.TypeInt8, Default: 0, SchemaType: map[string]string{"mysql": "tinyint unsigned"}},
		{Name: "floor", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "create_at", Type: field.TypeInt64, Default: 0, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
		{Name: "comment_replies", Type: field.TypeInt64, Nullable: true, SchemaType: map[string]string{"mysql": "int UNSIGNED"}},
	}
	// ReplyTable holds the schema information for the "reply" table.
	ReplyTable = &schema.Table{
		Name:       "reply",
		Columns:    ReplyColumns,
		PrimaryKey: []*schema.Column{ReplyColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reply_comment_replies",
				Columns:    []*schema.Column{ReplyColumns[9]},
				RefColumns: []*schema.Column{CommentColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "reply_user_id_post_id_comment_id",
				Unique:  false,
				Columns: []*schema.Column{ReplyColumns[1], ReplyColumns[2], ReplyColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommentTable,
		PostTable,
		ReplyTable,
	}
)

func init() {
	CommentTable.ForeignKeys[0].RefTable = PostTable
	CommentTable.Annotation = &entsql.Annotation{
		Table: "comment",
	}
	PostTable.Annotation = &entsql.Annotation{
		Table: "post",
	}
	ReplyTable.ForeignKeys[0].RefTable = CommentTable
	ReplyTable.Annotation = &entsql.Annotation{
		Table: "reply",
	}
}