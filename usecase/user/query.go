package user

import (
	"MyAPI/delivery/helper"
	"MyAPI/entity"
	"fmt"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

// GetList query for select all data
func GetList(db *xorm.Engine) (user []entity.User, err error) {
	err = db.Find(&user)
	return
}

// GetById query for select data by ID
func GetById(db *xorm.Engine, id int64) (user entity.User, err error) {
	_, err = db.Where(fmt.Sprintf("id=%d", id)).Get(&user)
	return
}

// GetTokenTable query for select data by token
func GetByToken(db *xorm.Engine, token string) (user entity.User, err error) {
	var tb_token entity.Token
	_, err = db.Where(fmt.Sprintf("token='%s'", token)).Get(&tb_token)
	if err != nil {
		return
	}

	_, err = db.Where(fmt.Sprintf("id=%d", tb_token.Id_user)).Get(&user)
	return
}

// GetToken query for generate token
func GetToken(db *xorm.Engine, logBody entity.LoginBody) (token entity.LoginResponse, err error) {
	var user entity.User
	_, err = db.Where(fmt.Sprintf("email='%s'", logBody.Email)).Get(&user)
	if err != nil {
		return
	}

	err = user.CheckPassword(logBody.Password)
	if err != nil || user.Status == "online" {
		return
	}

	jwtWrapper := helper.JwtWrapper{
		SecretKey:       "r4h4s14d0nk",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		return
	}

	token = entity.LoginResponse{
		Token: signedToken,
	}

	_, err = db.Where(fmt.Sprintf("id=%d", user.Id)).Update(
		entity.User{
			Status: "online",
		},
	)
	if err != nil {
		return
	}

	err = SetToken(db, entity.Token{
		Id_user: user.Id,
		Token:   token.Token,
	})
	if err != nil {
		return
	}

	return
}

// PostToken query for insert (initial) token
func PostToken(db *xorm.Engine, token entity.Token) (err error) {
	_, err = db.Insert(
		entity.Token{
			Id:      token.Id,
			Id_user: token.Id_user,
			Token:   "",
		},
	)
	return
}

// SetToken query for update token
func SetToken(db *xorm.Engine, token entity.Token) (err error) {
	_, err = db.Where(fmt.Sprintf("id_user=%d", token.Id_user)).Update(
		entity.Token{
			Token: token.Token,
		},
	)
	return
}

// DeleteToken query for delete token data
func DeleteToken(db *xorm.Engine, id int64) (err error) {
	var token entity.Token
	_, err = db.Where(fmt.Sprintf("id_user=%d", id)).Delete(&token)
	return
}

// Post query for insert data
func Post(db *xorm.Engine, user entity.User) (err error) {
	_, err = db.Insert(
		entity.User{
			Id:       user.Id,
			Email:    user.Email,
			Password: user.Password,
			Username: user.Username,
			Status:   "offline",
		},
	)
	return
}

// Put query for update data
func Put(db *xorm.Engine, user entity.User) (err error) {
	_, err = db.Where(fmt.Sprintf("id=%d", user.Id)).Update(
		entity.User{
			Email:    user.Email,
			Password: user.Password,
			Username: user.Username,
			Status:   user.Status,
		},
	)
	return
}

// Delete query for delete data
func Delete(db *xorm.Engine, id int64) (err error) {
	var user entity.User
	err = DeleteToken(db, id)
	if err != nil {
		return
	}
	_, err = db.Where(fmt.Sprintf("id=%d", id)).Delete(&user)
	return
}
