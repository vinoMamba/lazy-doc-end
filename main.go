package main

import (
	"github.com/vinoMamba/lazy-doc-end/handler"
	"github.com/vinoMamba/lazy-doc-end/server"
)

func main() {
	r := server.SetupHttpServer()
	handler.HandleUser(r)
	r.Run(":3000")
}
