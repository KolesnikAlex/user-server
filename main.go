package main

import (
	"github.com/KolesnikAlex/user-server/app"
	"github.com/KolesnikAlex/user-server/http"
)

func main() {
	app := app.New()

	http.Setup()

	app.Run()
}
