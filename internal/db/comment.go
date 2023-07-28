package db

//
//import (
//	"context"
//	"database/sql"
//	"go-rest/internal/comment"
//)
//
//type CommentRow struct {
//	ID     string
//	Slug   sql.NullString
//	Body   sql.NullString
//	Author sql.NullString
//}
//
//func (d *Database) GetComment(
//	ctx context.Context,
//	uuid string,
//) (comment.Comment, error) {
//	var cmtRow CommentRow
//	row := d.Client.QueryRowContext(
//		ctx,
//		`SELECT id, slug,body, author FROM comments WHERE id =$1`,
//		uuid,
//	)
//	return comment.Comment{}, nil
//}
