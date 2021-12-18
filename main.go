package main

import "user-server/app"

func main() {
	app := app.New()

	//http.Setup()

	app.Run()
}
