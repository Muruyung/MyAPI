package todo

import (
	"MyAPI/config"
	"MyAPI/controller/todo"
	"MyAPI/entity"
	"fmt"

	"github.com/kataras/iris/v12"
	"xorm.io/xorm"
)

func Init(app *iris.Application, db *xorm.Engine) {
	err := db.Sync2(new(entity.Todo))

	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Sync2(new(entity.Token))

	if err != nil {
		fmt.Println(err)
		return
	}

	userAPI := config.Party(app, "/todo")
	{
		userAPI.Use(iris.Compression)
		userAPI.Post("/", todo.Init(db).Create)
		userAPI.Get("/", todo.Init(db).ReadList)
		userAPI.Get("/{id}", todo.Init(db).ReadById)
		userAPI.Put("/", todo.Init(db).Update)
		userAPI.Delete("/{id}", todo.Init(db).Delete)
	}
}
