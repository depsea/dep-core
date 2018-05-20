package main

import (
	"fmt"

	"github.com/depsea/dep-core/app"
	"github.com/depsea/dep-core/core"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := core.NewDB()
	defer db.Close()

	app.NewRouter(r)

	r.Run(fmt.Sprintf(":%d", core.PORT))
}
