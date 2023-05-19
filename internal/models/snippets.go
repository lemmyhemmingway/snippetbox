package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	// postgres uses $1
	// mysql uses ?
	sqlStatement := `
    INSERT INTO snippets (title, content, expires, created)
    VALUES ($1, $2, $3, $4)
    RETURNING id`
	now := time.Now()
	exp := now.Add(time.Duration(expires))
	id := 0
	err := m.DB.QueryRow(sqlStatement, title, content, exp, now).Scan(&id)
	if err != nil {
		panic(err)
	}

	return id, nil
}
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	sqlStatement := `SELECT * FROM snippets
    WHERE id = $1`

	s := &Snippet{}
	err := m.DB.QueryRow(sqlStatement, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	sqlStatement := `SELECT * FROM snippets
    ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {
		s := &Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
