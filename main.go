package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/vinoMamba/lazy-doc-end/bootstrap"
	btsConfig "github.com/vinoMamba/lazy-doc-end/config"
	"github.com/vinoMamba/lazy-doc-end/pkg/config"
)

func init() {
	btsConfig.Initialize()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "load the .env file")
	flag.Parse()
	config.InitConfig(env)

	port := config.Get("app.port")
	fmt.Println(port)

	r := gin.New()
	bootstrap.SetupRoute(r)

	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
