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
	sql := `SELECT t1.* FROM cw_post t1, 
	(SELECT id FROM cw_post WHERE status=? ORDER BY last_update DESC, id DESC LIMIT ?,?) t2 
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

func InsertReply(ctx context.Context, c *Client, userID, postID, commentID, parentID int64, content string) (int64, error) {
	db := c.driver.(*sql.Driver).DB()
	//先获取楼层数
	var floor int64
	// 行锁
	sql := "SELECT count(*) AS total FROM cw_reply WHERE comment_id=? FOR UPDATE"
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	if err := db.QueryRowContext(ctx, sql, commentID).Scan(&floor); err != nil {
		tx.Rollback()
		return 0, err
	}

	// 再插入
	r, err := c.Reply.Create().
		SetUserID(userID).
		SetPostID(postID).
		SetCommentID(commentID).
		SetParentID(parentID).
		SetFloor(floor + 1).
		SetContent(content).
		SetStatus(0).
		SetCreateAt(time.Now().Unix()).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return r.ID, nil
}
