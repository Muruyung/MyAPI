package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64  `xorm:"not null 'id' unique" json:"id"`
	Username string `xorm:"not null 'username' unique" json:"username"`
	Email    string `xorm:"not null 'email' unique" json:"email"`
	Password string `xorm:"not null 'password'" json:"password"`
	Status   string `xorm:"not null 'status'" json:"status"`
}

// NewUser create new user
func NewUser(username string, email string, password string) (user User) {
	id := time.Now().UnixNano() / int64(time.Millisecond)

	user = User{
		Id:       id,
		Username: username,
		Email:    email,
		Password: password,
		Status:   "offline",
	}

	return
}

// HashPassword encrypts user password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

// CheckPassword checks user password
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
