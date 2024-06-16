package main

import "task-tracker/app/db"

type App struct {
	db *db.DB
}

func createApp() *App {
	return &App{
		db: db.New(),
	}
}

func main() {
	app := createApp()
	app.db.Init()
	app.InitFlags()
}
