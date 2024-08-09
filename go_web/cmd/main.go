package main

import "git_truongvudinh/go_web/internal/app"

func main() {
	engine := app.NewGinEngine()
	engine.Run(":8080")
}
