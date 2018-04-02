package main

import (
	"fmt"

	"github.com/depsea/api/app"
	"github.com/depsea/api/core"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	core.NewDB()
	defer core.DB.Close()

	app.NewRouter(r)

	r.Run(fmt.Sprintf(":%d", core.PORT))
}
