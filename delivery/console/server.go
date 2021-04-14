package console

import (
	"MyAPI/config"
	"MyAPI/delivery/http/todo"
	"MyAPI/delivery/http/user"
)

func Run() {
	app := config.New()
	db := config.DBase()
	user.Init(app, db)
	todo.Init(app, db)
	config.Listen(app)
}
