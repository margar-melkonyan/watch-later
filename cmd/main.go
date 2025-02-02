package main

import (
	"github.com/margar-melkonyan/watch-later.git/internal/app"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		app.RunMigrations()
		return
	}
	app.RunApplication()
}
