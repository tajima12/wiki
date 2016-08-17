package model

import "database/sql"

// CommentAll ...
func CommentAll(db *sql.DB, id int) ([]Comment, error) {
	rows, err := db.Query(`select * from comments where article_id = ?`, id)
	if err != nil {
		return nil, err
	}
	return ScanComments(rows)
}

// NewComment ...
func NewComment(tx *sql.Tx, id int, body, user string) (sql.Result, error) {
	stmt, err := tx.Prepare(`
	insert into comments(article_id, body, user)
	values(?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id, body, user)
}
