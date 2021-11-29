package ent

import (
	"context"

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
