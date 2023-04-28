package main

import (
	"github.com/KolesnikAlex/user-server/app"
	"github.com/KolesnikAlex/user-server/http"
)

func main() {
	app := app.New()

	http.Setup()

	go app.RunHTTP()

	go app.RunGRPC()

	<-app.ServerOk
	<-app.ServerOk
	// main
	// main 2
}
