package main

import (
	"github.com/depsea/dep-core/config"
	"github.com/depsea/dep-core/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.Register(r)

	config, err := config.GetConfig()

	if err != nil {
		panic(err)
	}

	r.Run(config.Server.Addr)
}
