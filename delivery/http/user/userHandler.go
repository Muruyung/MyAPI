package user

import (
	"MyAPI/config"
	"MyAPI/controller/user"
	"MyAPI/entity"
	"fmt"

	"github.com/kataras/iris/v12"
	"xorm.io/xorm"
)

func Init(app *iris.Application, db *xorm.Engine) {
	err := db.Sync2(new(entity.User))

	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Sync2(new(entity.Token))

	if err != nil {
		fmt.Println(err)
		return
	}

	userAPI := config.Party(app, "/user")
	{
		userAPI.Use(iris.Compression)
		userAPI.Post("/", user.Init(db).Create)
		userAPI.Get("/", user.Init(db).ReadList)
		userAPI.Get("/{id}", user.Init(db).ReadById)
		userAPI.Put("/", user.Init(db).Update)
		userAPI.Delete("/{id}", user.Init(db).Delete)
	}

	logInAPI := config.Party(app, "/login")
	{
		logInAPI.Use(iris.Compression)
		logInAPI.Post("/", user.Init(db).Login)
	}

	logOutAPI := config.Party(app, "/logout")
	{
		logOutAPI.Use(iris.Compression)
		logOutAPI.Post("/", user.Init(db).Logout)
	}

	signUpAPI := config.Party(app, "/signup")
	{
		signUpAPI.Use(iris.Compression)
		signUpAPI.Post("/", user.Init(db).Create)
	}
}
