package main

import "flag"

func (app App) InitFlags() {
	var start string
	var complete bool
	flag.StringVar(&start, "start", "", "start a new task")
	flag.BoolVar(&complete, "complete", false, "complete the current task")
	flag.Parse()
	if start != "" {
		app.db.AddNewTask(start)
	}
	if complete {
		app.db.CompleteTask()
	}
}
