package main

import (
	"flag"
	"fmt"
	"time"
)

func (app App) InitFlags() {
	var start string
	var complete bool
	var total bool
  var export bool
	var month, year int
  var tag string
	flag.StringVar(&start, "start", "", "start a new task")
	flag.IntVar(&month, "month", int(time.Now().Month()), "month for query")
	flag.IntVar(&year, "year", int(time.Now().Year()), "year for query")
	flag.BoolVar(&complete, "complete", false, "complete the current task")
  flag.StringVar(&tag, "tag", "work", "tag for work, personal. default tag: work. all tag only should be used when querying")
	flag.BoolVar(&total, "total", false, "show total for month")
  flag.BoolVar(&export, "export", false, "export task to csv")
	flag.Parse()
  if tag != "work" && tag != "personal" && tag != "all" {
    fmt.Println("Invalid tag")
    return
  }
	if start != "" {
    if tag == "all" {
      fmt.Println("all tag should not be used when creating new task")
    }
		app.db.AddNewTask(start, tag)
	}
	if complete {
		app.db.CompleteTask()
	}
	if total {
		app.db.TotalForMonth(month, year, tag)
	}
  if export {
    app.db.Export(month, year, tag)
  }
}
