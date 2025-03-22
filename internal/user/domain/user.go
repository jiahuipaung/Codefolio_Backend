package domain

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint64
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(username, password, email string) (*User, error) {
	if username == "" || password == "" || email == "" {
		return nil, errors.New("username, password and email are required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &User{
		Username:  username,
		Password:  string(hashedPassword),
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
