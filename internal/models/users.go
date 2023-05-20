package models

import (
	"database/sql"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	now := time.Now()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	sqlStatement := `INSERT INTO users(name, email, hashed_password, created)
    VALUES($1, $2, $3, $4) RETURNING id`
	id := 0

	err = m.DB.QueryRow(sqlStatement, name, email, string(hashedPassword), now).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	sqlStatement := `SELECT id, hashed_password from users WHERE email = $1`
	err := m.DB.QueryRow(sqlStatement, email).Scan(&id, &hashedPassword)
	if err != nil {
		panic(err)
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		panic(err)
	}

	return id, nil
}
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
