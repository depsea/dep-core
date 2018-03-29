package main

import (
	"fmt"

	"github.com/depsea/api/core"
)

func main() {
	core.NewEngine().Run(fmt.Sprintf(":%d", core.PORT))
}
