package db

import "log"

type Project struct {
  id int
  name string
  tag string
}

func (db DB) createProjectTable() {
	q := "CREATE TABLE IF NOT EXISTS projects(id SERIAL, name TEXT, tag TEXT)"
	conn := db.getConnection()
	defer conn.Close()
	_, err := conn.Exec(q)
	if err != nil {
		log.Fatalln(err)
	}
}

func (db DB) addNewProject(name string, tag string) {
  q := "INSERT INTO projects(name, tag) VALUE($1, $2)"
  conn := db.getConnection()
  defer conn.Close()
  _, err := conn.Exec(q, name, tag)
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

