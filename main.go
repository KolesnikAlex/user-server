package main

import (
	"user-server/app"
	"user-server/http"
)

func main() {
	app := app.New()

	http.Setup()

	app.Run()
}
