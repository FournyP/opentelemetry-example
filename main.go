package main

import (
	"context"
)

func main() {
	app := InitializeApp()
	defer app.Cleanup(context.Background())

	app.FiberApp.Listen(":8000")
}
