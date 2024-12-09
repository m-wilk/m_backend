package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRole string

const (
	SuperAdminRole UserRole = "super-admin"
	AdminRole      UserRole = "admin"
	ClientRole     UserRole = "client"
)

type User struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string    `bson:"username" json:"username"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"-"`
	Role      UserRole  `bson:"role" json:"role"`
	CreatedAt time.Time `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time `bson:"updated_at" json:"updatedAt"`

	// valid password
	isPasswordHashed bool
}

func (u *User) IsValid() bool {
	return u.isPasswordHashed
}

func (u *User) HashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return "", err
	}

	stringHash := string(bytes)
	u.Password = stringHash
	u.isPasswordHashed = true
	return stringHash, nil
}

func (u *User) IsValidPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
