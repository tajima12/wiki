package model

import "database/sql"

func CommentAll(db *sql.DB, id int) ([]Comment, error) {
	rows, err := db.Query(`select * from comments where article_id = ?`, id)
	if err != nil {
		return nil, err
	}
	return ScanComments(rows)
}

func NewComment(tx *sql.Tx, id int, body string) (sql.Result, error) {
	stmt, err := tx.Prepare(`
	insert into comments(article_id, body)
	values(?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id, body)
}
