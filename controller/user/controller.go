package user

import (
	"MyAPI/delivery/helper"
	"MyAPI/entity"
	"MyAPI/usecase/user"
	"fmt"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

// Login controller using post method
func (db Database) Login(ctx iris.Context) {
	var logBody entity.LoginBody

	err := ctx.ReadJSON(&logBody)
	if err != nil {
		ctx.StopWithStatus(iris.StatusBadRequest)
		helper.CreateErrorResponse(ctx, err).BadRequest().JSON()
		return
	}

	token, err := user.GetToken(db.engine, logBody)
	if err != nil {
		ctx.StopWithStatus(iris.StatusUnauthorized)
		helper.CreateErrorResponse(ctx, err).Unauthorized().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetData(token).JSON()
}

// Logout controller using post method
func (db Database) Logout(ctx iris.Context) {
	var logResponse entity.LoginResponse
	err := ctx.ReadJSON(&logResponse)

	if err != nil {
		ctx.StopWithStatus(iris.StatusBadRequest)
		helper.CreateErrorResponse(ctx, err).BadRequest().JSON()
		return
	}

	tb_user, err := user.GetByToken(db.engine, logResponse.Token)
	if err != nil {
		ctx.StopWithStatus(iris.StatusNotFound)
		helper.CreateErrorResponse(ctx, err).NotFound().JSON()
		return
	}

	tb_user.Status = "offline"
	err = user.Put(db.engine, tb_user)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	err = user.SetToken(db.engine, entity.Token{
		Id_user: tb_user.Id,
		Token:   "-",
	})
	if err != nil {
		fmt.Println("anjay")
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetData(tb_user).JSON()
}

// Create controller using post method
func (db Database) Create(ctx iris.Context) {
	var tb_user entity.User

	err := ctx.ReadJSON(&tb_user)
	if err != nil {
		ctx.StopWithStatus(iris.StatusBadRequest)
		helper.CreateErrorResponse(ctx, err).BadRequest().JSON()
		return
	}

	err = tb_user.HashPassword(tb_user.Password)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	tb_user.Id = time.Now().UnixNano() / int64(time.Millisecond)
	err = user.Post(db.engine, tb_user)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	var token entity.Token
	token.Id = time.Now().UnixNano() / int64(time.Millisecond)
	token.Id_user = tb_user.Id
	err = user.PostToken(db.engine, token)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetData(tb_user).JSON()
}

// ReadList controller using get method
func (db Database) ReadList(ctx iris.Context) {
	tb_user, err := user.GetList(db.engine)

	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetData(tb_user).JSON()
}

// ReadById controller using get method
func (db Database) ReadById(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)
	if err != nil {
		ctx.StopWithStatus(iris.StatusBadRequest)
		helper.CreateErrorResponse(ctx, err).BadRequest().JSON()
		return
	}

	tb_user, err := user.GetById(db.engine, id)

	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetData(tb_user).JSON()
}

// Update controller using put method
func (db Database) Update(ctx iris.Context) {
	var tb_user entity.User
	err := ctx.ReadJSON(&tb_user)

	if err != nil {
		ctx.StopWithStatus(iris.StatusBadRequest)
		helper.CreateErrorResponse(ctx, err).BadRequest().JSON()
		return
	}

	err = tb_user.HashPassword(tb_user.Password)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	err = user.Put(db.engine, tb_user)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetData(tb_user).JSON()
}

// Delete controller using delete method
func (db Database) Delete(ctx iris.Context) {
	id, err := strconv.ParseInt(ctx.Params().Get("id"), 10, 64)

	if err != nil {
		ctx.StopWithStatus(iris.StatusBadRequest)
		helper.CreateErrorResponse(ctx, err).BadRequest().JSON()
		return
	}

	err = user.Delete(db.engine, id)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		helper.CreateErrorResponse(ctx, err).InternalServer().JSON()
		return
	}

	helper.CreateResponse(ctx).Ok().SetMessage(fmt.Sprintf("%d deleted", id)).JSON()
}
