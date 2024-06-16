package db

import (
	"strings"
	"task-tracker/app/util"
	"time"
)

type taskWithFilter struct {
	s_no       int
	message    string
	day        int
	month      time.Month
	year       int
	date       string
	start_time string
	stop_time  string
	total      int
	tag        string
}

func (db DB) getTotalForMonth(month int, year int, tag string) (total int) {
	tasks := db.filterMonthYear(month, year, tag)
	total = 0
	for _, task := range tasks {
		total += task.total
	}
	return
}

func (db DB) filterMonthYear(month int, year int, tag string) (result []taskWithFilter) {
	tasks := convertTaskToFiltered(db.getTasks())
	for _, task := range tasks {
		if time.Month(month) == task.month && year == task.year && (task.tag == tag || tag == "all") {
			result = append(result, task)
		}
	}
	return
}

func filterTag(tasks []taskWithFilter, tag string) (result []taskWithFilter) {
	if tag == "all" {
		return tasks
	}
	for _, task := range tasks {
		if task.tag == tag {
			result = append(result, task)
		}
	}
	return
}

func convertTaskToFiltered(tasks []Task) (result []taskWithFilter) {
	for i, task := range tasks {
		newTask := taskWithFilter{
			s_no:       i + 1,
			message:    task.message,
			date:       task.date[:10],
			start_time: task.start_time[11:19],
			stop_time:  task.stop_time[11:19],
			month:      time.Month(util.StringToInt(strings.Split(task.date[:10], "-")[1])),
			year:       util.StringToInt(task.date[:4]),
			day:        util.StringToInt(task.date[8:10]),
			total:      task.total,
			tag:        task.tag,
		}
		result = append(result, newTask)
	}
	return
}

func filterDay(tasks []taskWithFilter, day int, month int, year int) (result []taskWithFilter) {
	for _, task := range tasks {
		if time.Month(month) == task.month && year == task.year && day == task.day {
      result = append(result, task)
		}
	}
	return
}

