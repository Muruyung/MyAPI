package todo

import (
	"MyAPI/delivery/helper"
	"MyAPI/entity"
	"MyAPI/usecase/todo"
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

// Create controller using post method
func (db Database) Create(ctx iris.Context) {
	var todos entity.Todo

	err := ctx.ReadJSON(&todos)
	if err != nil {
		ctx.StopWithStatus(iris.StatusBadRequest)
		helper.CreateErrorResponse(ctx, err).BadRequest().JSON()
		return
	}

	todos.Id = time.Now().UnixNano() / int64(time.Millisecond)
	err = todo.Post(db.engine, todos)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetData(todos).JSON()
}

// ReadList controller using get method
func (db Database) ReadList(ctx iris.Context) {
	todos, err := todo.GetList(db.engine)

	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetData(todos).JSON()
}

// ReadById controller using get method
func (db Database) ReadById(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
	if err != nil {
		ctx.StopWithStatus(iris.StatusBadRequest)
		helper.CreateErrorResponse(ctx, err).BadRequest().JSON()
		return
	}

	todos, err := todo.GetById(db.engine, id)

	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetData(todos).JSON()
}

// Update controller using put method
func (db Database) Update(ctx iris.Context) {
	var todos entity.Todo
	err := ctx.ReadJSON(&todos)

	if err != nil {
		ctx.StopWithStatus(iris.StatusBadRequest)
		helper.CreateErrorResponse(ctx, err).BadRequest().JSON()
		return
	}

	err = todo.Put(db.engine, todos)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetData(todos).JSON()
}

// Delete controller using delete method
func (db Database) Delete(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	if err != nil {
		ctx.StopWithStatus(iris.StatusBadRequest)
		helper.CreateErrorResponse(ctx, err).BadRequest().JSON()
		return
	}

	err = todo.Delete(db.engine, id)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetMessage(fmt.Sprintf("%d deleted", id)).JSON()
}
