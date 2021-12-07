package main

func main() {
	app := app.New()

	http.Setup()

	app.Run()
}
