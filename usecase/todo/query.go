package todo

import (
	"MyAPI/entity"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

// GetList query for select all data
func GetList(db *xorm.Engine) (todo []entity.Todo, err error) {
	err = db.Where("NOT id_user=-1").Find(&todo)
	return
}

// GetById query for select data by ID
func GetById(db *xorm.Engine, id int64) (todo entity.Todo, err error) {
	_, err = db.Where(fmt.Sprintf("id=%d", id)).Get(&todo)
	return
}

// Post query for insert data
func Post(db *xorm.Engine, todo entity.Todo) (err error) {
	_, err = db.Insert(
		entity.Todo{
			Id:          todo.Id,
			Id_user:     todo.Id_user,
			Name:        todo.Name,
			Description: todo.Description,
		},
	)
	return
}

// Put query for update data
func Put(db *xorm.Engine, todo entity.Todo) (err error) {
	_, err = db.Where(fmt.Sprintf("id=%d", todo.Id)).Update(
		entity.Todo{
			Id_user:     todo.Id_user,
			Name:        todo.Name,
			Description: todo.Description,
		},
	)
	return
}

// Delete query for delete data
func Delete(db *xorm.Engine, id int64) (err error) {
	_, err = db.Where(fmt.Sprintf("id=%d", id)).Update(
		entity.Todo{
			Id_user:    -1,
			Deleted_at: time.Now(),
		},
	)
	return
}
