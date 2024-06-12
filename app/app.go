package main

import "task-tracker/app/db"

type App struct {
	db *db.DB
}

func createApp() *App {
	return &App{
		db: db.Connect(),
	}
}

func main() {
	app := createApp()
	app.InitFlags()
	app.db.Init()
}
