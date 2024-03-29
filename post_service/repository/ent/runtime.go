// Code generated by ent, DO NOT EDIT.

package ent

import (
	"cw_post_service/repository/ent/comment"
	"cw_post_service/repository/ent/post"
	"cw_post_service/repository/ent/reply"
	"cw_post_service/repository/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescUserID is the schema descriptor for user_id field.
	commentDescUserID := commentFields[1].Descriptor()
	// comment.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	comment.UserIDValidator = commentDescUserID.Validators[0].(func(uint64) error)
	// commentDescPostID is the schema descriptor for post_id field.
	commentDescPostID := commentFields[2].Descriptor()
	// comment.PostIDValidator is a validator for the "post_id" field. It is called by the builders before save.
	comment.PostIDValidator = commentDescPostID.Validators[0].(func(uint64) error)
	// commentDescContent is the schema descriptor for content field.
	commentDescContent := commentFields[3].Descriptor()
	// comment.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	comment.ContentValidator = commentDescContent.Validators[0].(func(string) error)
	// commentDescStatus is the schema descriptor for status field.
	commentDescStatus := commentFields[4].Descriptor()
	// comment.DefaultStatus holds the default value on creation for the status field.
	comment.DefaultStatus = commentDescStatus.Default.(int8)
	// comment.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	comment.StatusValidator = commentDescStatus.Validators[0].(func(int8) error)
	// commentDescFloor is the schema descriptor for floor field.
	commentDescFloor := commentFields[5].Descriptor()
	// comment.FloorValidator is a validator for the "floor" field. It is called by the builders before save.
	comment.FloorValidator = commentDescFloor.Validators[0].(func(uint64) error)
	// commentDescCreateAt is the schema descriptor for create_at field.
	commentDescCreateAt := commentFields[6].Descriptor()
	// comment.DefaultCreateAt holds the default value on creation for the create_at field.
	comment.DefaultCreateAt = commentDescCreateAt.Default.(int64)
	// comment.CreateAtValidator is a validator for the "create_at" field. It is called by the builders before save.
	comment.CreateAtValidator = commentDescCreateAt.Validators[0].(func(int64) error)
	// commentDescID is the schema descriptor for id field.
	commentDescID := commentFields[0].Descriptor()
	// comment.IDValidator is a validator for the "id" field. It is called by the builders before save.
	comment.IDValidator = commentDescID.Validators[0].(func(uint64) error)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescUserID is the schema descriptor for user_id field.
	postDescUserID := postFields[1].Descriptor()
	// post.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	post.UserIDValidator = postDescUserID.Validators[0].(func(uint64) error)
	// postDescTitle is the schema descriptor for title field.
	postDescTitle := postFields[2].Descriptor()
	// post.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	post.TitleValidator = postDescTitle.Validators[0].(func(string) error)
	// postDescContent is the schema descriptor for content field.
	postDescContent := postFields[3].Descriptor()
	// post.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	post.ContentValidator = postDescContent.Validators[0].(func(string) error)
	// postDescStatus is the schema descriptor for status field.
	postDescStatus := postFields[4].Descriptor()
	// post.DefaultStatus holds the default value on creation for the status field.
	post.DefaultStatus = postDescStatus.Default.(int8)
	// post.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	post.StatusValidator = postDescStatus.Validators[0].(func(int8) error)
	// postDescReplyNum is the schema descriptor for reply_num field.
	postDescReplyNum := postFields[5].Descriptor()
	// post.DefaultReplyNum holds the default value on creation for the reply_num field.
	post.DefaultReplyNum = postDescReplyNum.Default.(int64)
	// post.ReplyNumValidator is a validator for the "reply_num" field. It is called by the builders before save.
	post.ReplyNumValidator = postDescReplyNum.Validators[0].(func(int64) error)
	// postDescCreateAt is the schema descriptor for create_at field.
	postDescCreateAt := postFields[6].Descriptor()
	// post.DefaultCreateAt holds the default value on creation for the create_at field.
	post.DefaultCreateAt = postDescCreateAt.Default.(int64)
	// post.CreateAtValidator is a validator for the "create_at" field. It is called by the builders before save.
	post.CreateAtValidator = postDescCreateAt.Validators[0].(func(int64) error)
	// postDescUpdateAt is the schema descriptor for update_at field.
	postDescUpdateAt := postFields[7].Descriptor()
	// post.DefaultUpdateAt holds the default value on creation for the update_at field.
	post.DefaultUpdateAt = postDescUpdateAt.Default.(int64)
	// post.UpdateAtValidator is a validator for the "update_at" field. It is called by the builders before save.
	post.UpdateAtValidator = postDescUpdateAt.Validators[0].(func(int64) error)
	// postDescPin is the schema descriptor for pin field.
	postDescPin := postFields[8].Descriptor()
	// post.DefaultPin holds the default value on creation for the pin field.
	post.DefaultPin = postDescPin.Default.(int8)
	// post.PinValidator is a validator for the "pin" field. It is called by the builders before save.
	post.PinValidator = postDescPin.Validators[0].(func(int8) error)
	// postDescID is the schema descriptor for id field.
	postDescID := postFields[0].Descriptor()
	// post.IDValidator is a validator for the "id" field. It is called by the builders before save.
	post.IDValidator = postDescID.Validators[0].(func(uint64) error)
	replyFields := schema.Reply{}.Fields()
	_ = replyFields
	// replyDescUserID is the schema descriptor for user_id field.
	replyDescUserID := replyFields[1].Descriptor()
	// reply.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	reply.UserIDValidator = replyDescUserID.Validators[0].(func(uint64) error)
	// replyDescCommentID is the schema descriptor for comment_id field.
	replyDescCommentID := replyFields[2].Descriptor()
	// reply.CommentIDValidator is a validator for the "comment_id" field. It is called by the builders before save.
	reply.CommentIDValidator = replyDescCommentID.Validators[0].(func(uint64) error)
	// replyDescParentID is the schema descriptor for parent_id field.
	replyDescParentID := replyFields[3].Descriptor()
	// reply.ParentIDValidator is a validator for the "parent_id" field. It is called by the builders before save.
	reply.ParentIDValidator = replyDescParentID.Validators[0].(func(uint64) error)
	// replyDescContent is the schema descriptor for content field.
	replyDescContent := replyFields[4].Descriptor()
	// reply.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	reply.ContentValidator = replyDescContent.Validators[0].(func(string) error)
	// replyDescStatus is the schema descriptor for status field.
	replyDescStatus := replyFields[5].Descriptor()
	// reply.DefaultStatus holds the default value on creation for the status field.
	reply.DefaultStatus = replyDescStatus.Default.(int8)
	// reply.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	reply.StatusValidator = replyDescStatus.Validators[0].(func(int8) error)
	// replyDescFloor is the schema descriptor for floor field.
	replyDescFloor := replyFields[6].Descriptor()
	// reply.FloorValidator is a validator for the "floor" field. It is called by the builders before save.
	reply.FloorValidator = replyDescFloor.Validators[0].(func(uint64) error)
	// replyDescCreateAt is the schema descriptor for create_at field.
	replyDescCreateAt := replyFields[7].Descriptor()
	// reply.DefaultCreateAt holds the default value on creation for the create_at field.
	reply.DefaultCreateAt = replyDescCreateAt.Default.(int64)
	// reply.CreateAtValidator is a validator for the "create_at" field. It is called by the builders before save.
	reply.CreateAtValidator = replyDescCreateAt.Validators[0].(func(int64) error)
	// replyDescID is the schema descriptor for id field.
	replyDescID := replyFields[0].Descriptor()
	// reply.IDValidator is a validator for the "id" field. It is called by the builders before save.
	reply.IDValidator = replyDescID.Validators[0].(func(uint64) error)
}
