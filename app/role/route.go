package role

import (
	"github.com/depsea/api/core"
	"github.com/depsea/api/middlewares"
	"github.com/gin-gonic/gin"
)

// RoleModel -
var RoleModel = NewModel()

// New -
func New(e *gin.Engine) {
	e.GET("/roles", list)
	e.GET("/roles/:role_id", detail)
	e.POST("/roles", middlewares.Auth(), create)
	e.PUT("/roles/:role_id", middlewares.Auth(), update)
	e.DELETE("/roles/:role_id", middlewares.Auth(), delete)
}

func list(ctx *gin.Context) {
	roles, err := RoleModel.Find()

	if err != nil {
		core.DBError(ctx, []Role{})
		return
	}

	core.Result(ctx, roles)
}

func detail(ctx *gin.Context) {
	_id := core.ParseIDParam(ctx, "role_id")

	role, err := RoleModel.FindOne(_id)

	if err != nil {
		core.DBError(ctx)
		return
	}

	core.Result(ctx, role)
}

func create(ctx *gin.Context) {
	var role *Role
	err := ctx.BindJSON(&role)

	if err != nil {
		core.BodyError(ctx)
		return
	}

	err = RoleModel.Create(role)

	if err != nil {
		core.DBError(ctx)
		return
	}

	core.Result(ctx, *role)
}

func update(ctx *gin.Context) {
	_id := core.ParseIDParam(ctx, "role_id")

	var role *Role
	err := ctx.BindJSON(&role)

	if err != nil {
		core.BodyError(ctx)
		return
	}
	err = RoleModel.Update(_id, role)

	if err != nil {
		core.DBError(ctx)
		return
	}
	core.Result(ctx, *role)
}

func delete(ctx *gin.Context) {
	_id := core.ParseIDParam(ctx, "role_id")

	role, err := RoleModel.Delete(_id)

	if err != nil {
		core.DBError(ctx)
		return
	}

	core.Result(ctx, role)
}
