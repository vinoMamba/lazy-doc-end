package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/vinoMamba/lazy-doc-end/bootstrap"
)

func main() {
	r := gin.New()
	bootstrap.SetupRoute(r)

	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
