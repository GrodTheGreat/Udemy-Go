package models

import (
	"errors"
	"rest/db"
	"rest/utils"
)

type User struct {
	ID           int64  `json:"id"`
	Email        string `binding:"required" json:"email"`
	PasswordHash string `binding:"required" json:""`
}

type SignupRequest struct {
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (user *User) Save(password string) error {
	query := `
INSERT INTO user (email, password_hash)
VALUES (?, ?)
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer utils.CheckClose(stmt, &err)

	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hash)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = userId

	return nil
}

func (user *User) ValidateCredentials(password string) error {
	query := `SELECT id, password_hash FROM user WHERE email = ? LIMIT 1`
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPasswordHash string
	err := row.Scan(&user.ID, &retrievedPasswordHash)
	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(retrievedPasswordHash, password)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	user.PasswordHash = retrievedPasswordHash

	return nil
}
