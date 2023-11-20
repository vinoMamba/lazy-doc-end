package main

import (
	"github.com/vinoMamba/lazy-doc-end/config"
	"github.com/vinoMamba/lazy-doc-end/handler"
	"github.com/vinoMamba/lazy-doc-end/server"
	"github.com/vinoMamba/lazy-doc-end/storage"
)

func init() {
	config.LoadConfig(".")
	storage.DbConn()
}

func main() {
	r := server.SetupHttpServer()
	handler.HandleUser(r)
	r.Run(":3000")
}
