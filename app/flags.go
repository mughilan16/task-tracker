package main

import (
	"flag"
	"time"
)

func (app App) InitFlags() {
	var start string
	var complete bool
	var total bool
  var export bool
	var month, year int
	flag.StringVar(&start, "start", "", "start a new task")
	flag.IntVar(&month, "month", int(time.Now().Month()), "month for query")
	flag.IntVar(&year, "year", int(time.Now().Year()), "year for query")
	flag.BoolVar(&complete, "complete", false, "complete the current task")
	flag.BoolVar(&total, "total", false, "show total for month")
  flag.BoolVar(&export, "export", false, "export task to csv")
	flag.Parse()
	if start != "" {
		app.db.AddNewTask(start)
	}
	if complete {
		app.db.CompleteTask()
	}
	if total {
		app.db.TotalForMonth(month, year)
	}
  if export {
    app.db.Export(month, year)
  }
}
