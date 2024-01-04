package main

import (
	"os"

	"github.com/vinoMamba/lazydoc/internal/lazydoc"
)

func main() {
	cmd := lazydoc.NewRootCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
