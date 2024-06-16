package db

import "log"

type Project struct {
  id int
  name string
  tag string
}

func (db DB) createProjectTable() {
	q := "CREATE TABLE IF NOT EXISTS projects(id SERIAL, name string, tag string)"
	conn := db.getConnection()
	defer conn.Close()
	_, err := conn.Exec(q)
	if err != nil {
		log.Fatalln(err)
	}
}

func (db DB) getProjects() (projects []Project){
	q := "SELECT (id, name, tag) FROM projects"
	conn := db.getConnection()
	defer conn.Close()
	rows, err := conn.Query(q)
	if err != nil {
		log.Fatalln(err)
	}
  for rows.Next() {
    var project Project
    rows.Scan(&project.id, &project.name, &project.tag)
  }
  return
}
