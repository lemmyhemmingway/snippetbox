package models

import (
	"database/sql"
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

func (m *SnippetModel) Insert (title string, content string, expires int) (int, error) {
    smt := `INSTERT INTO snippets (title, content, expires, created)
    VALUES (?, ?, '2022-02-02', '2022-02-02')`
    result, err := m.DB.Exec(smt, title, content, expires)
    if err != nil {
        return 0, err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return int(id), nil
}
func (m *SnippetModel) Get (id int) (*Snippet, error) {
   return nil, nil 
}
func (m *SnippetModel) Latest() ([]*Snippet, error) {
   return nil, nil 
}
