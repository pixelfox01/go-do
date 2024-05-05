package models

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
	Created   time.Time
}

type TodoModel struct {
	DB *sql.DB
}

func (m *TodoModel) Insert(title string, completed bool) (int, error) {
	query := `INSERT INTO todos (title, completed, created)
  VALUES (?, ?, UTC_TIMESTAMP)`

	result, err := m.DB.Exec(query, title, completed)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *TodoModel) Delete(id int) (int, error) {
	query := `DELETE FROM todos WHERE id=?`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rows), nil
}
