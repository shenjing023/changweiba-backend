package ent

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
)

// GetPosts get posts by page and page_size
// TODO 待优化
func GetPosts(ctx context.Context, c *Client, page, pageSize int) ([]*Post, error) {
	db := c.driver.(*sql.Driver).DB()
	sql := `SELECT t1.* FROM post t1, 
	(SELECT id FROM post WHERE status=? ORDER BY update_at DESC, id DESC LIMIT ?,?) t2 
	WHERE t1.id=t2.id`
	rows, err := db.QueryContext(ctx, sql, 0, pageSize*(page-1), pageSize)
	if err != nil {
		return nil, err
	}
	var posts []*Post
	defer rows.Close()
	for rows.Next() {
		var p Post
		rows.Scan(&p)
		posts = append(posts, &p)
	}
	return posts, nil
}

func InsertReply(ctx context.Context, c *Client, userID, commentID, parentID uint64, content string) (int64, error) {
	db := c.driver.(*sql.Driver).DB()
	//先获取楼层数
	var floor int64
	// 行锁
	sql := "SELECT count(*) AS total FROM reply WHERE comment_id=? FOR UPDATE"
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	if err := db.QueryRowContext(ctx, sql, commentID).Scan(&floor); err != nil {
		tx.Rollback()
		return 0, err
	}

	// 再插入
	rc := c.Reply.Create().
		SetUserID(userID).
		SetFloor(uint64(floor + 1)).
		SetContent(content).
		SetStatus(0).
		SetCreateAt(time.Now().Unix()).
		SetOwnerID(commentID)
	if parentID != 0 {
		rc.SetParentID(parentID)
	}
	r, err := rc.Save(ctx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return int64(r.ID), nil
}

// GetCommentsByPostID 获取帖子所属的评论
// TODO 待优化
func GetCommentsByPostID(ctx context.Context, c *Client, postID int64, page, pageSize int) ([]*Comment, error) {
	db := c.driver.(*sql.Driver).DB()
	sql := `SELECT 
				t1.*
			FROM 
				comment t1, 
			(SELECT id FROM comment WHERE status=? AND post_id=? LIMIT ?,?) t2 
			WHERE t1.id=t2.id`
	rows, err := db.QueryContext(ctx, sql, 0, pageSize*(page-1), pageSize)
	if err != nil {
		return nil, err
	}
	var comments []*Comment
	defer rows.Close()
	for rows.Next() {
		var c Comment
		rows.Scan(&c)
		comments = append(comments, &c)
	}
	return comments, nil
}

// GetCommentsByPostID 获取帖子所属的评论
// TODO 待优化
func GetCommentsByPostID1(ctx context.Context, c *Client, postID int64, page, pageSize int) ([]*Comment, error) {
	db := c.driver.(*sql.Driver).DB()
	sql := `SELECT 
				t1.*
			FROM 
				comment t1, 
			(SELECT id FROM comment WHERE status=? AND post_id=? LIMIT ?,?) t2 
			WHERE t1.id=t2.id`
	rows, err := db.QueryContext(ctx, sql, 0, pageSize*(page-1), pageSize)
	if err != nil {
		return nil, err
	}
	var comments []*Comment
	defer rows.Close()
	for rows.Next() {
		var c Comment
		rows.Scan(&c)
		comments = append(comments, &c)
	}
	return comments, nil
}
