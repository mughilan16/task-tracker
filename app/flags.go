package main

import (
	"flag"
	"fmt"
	"time"
)

func (app App) InitFlags() {
	var start, complete, total, export, today, currentMonth, info, create, projects bool
	var message, tag string
	var month, year int
	flag.BoolVar(&start, "start", false, "start a new task")
	flag.BoolVar(&create, "create", false, "create a new project")
	flag.StringVar(&message, "m", "", "description of task")
	flag.IntVar(&month, "month", int(time.Now().Month()), "month for query")
	flag.IntVar(&year, "year", int(time.Now().Year()), "year for query")
	flag.BoolVar(&complete, "complete", false, "complete the current task")
	flag.StringVar(&tag, "tag", "all", "tag for work, personal. default tag: all. all tag only should be used when querying")
	flag.BoolVar(&total, "total", false, "show total for month")
	flag.BoolVar(&info, "info", false, "show info of current task")
	flag.BoolVar(&export, "export", false, "export task to csv")
  flag.BoolVar(&today, "today", false, "set today date for filter")
  flag.BoolVar(&currentMonth, "currentmonth", false, "set current month for filter")
  flag.BoolVar(&projects, "projects", false, "list the projects")
	flag.Parse()
	if tag != "work" && tag != "personal" && tag != "all" {
		fmt.Println("Invalid tag")
		return
	}
	if start {
		if tag == "all" {
			fmt.Println("all tag should not be used when creating new task")
			return
		}
		if message == "" {
			fmt.Println("description should be provided to create a new task")
			return
		}
		app.db.AddNewTask(message, tag)
		return
	}
	if complete {
		app.db.CompleteTask()
		return
	}
	if total {
    if today {
      app.db.TodayTotal(tag)
      return
    }
    if currentMonth {
      app.db.ThisMonthTotal(tag)
      return
    }
		app.db.TotalForMonth(month, year, tag)
		return
	}
	if export {
		app.db.Export(month, year, tag)
		return
	}
  if info {
    app.db.ActiveTask(tag)
  }
  if create {
  }
  if projects {
  }
}
