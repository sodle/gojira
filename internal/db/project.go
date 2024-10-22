package db

import "database/sql"

type Project struct {
	Key    string
	Title  string
	Issues []Issue
}

func CreateProject(key string, title string) (*Project, error) {
	db := InitDb()

	statement, err := db.Prepare("INSERT INTO project (key, title) VALUES (?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(key, title)
	if err != nil {
		return nil, err
	}

	return &Project{Key: key, Title: title}, nil
}

func GetProjectByKey(key string) (*Project, error) {
	db := InitDb()

	statement, err := db.Prepare("SELECT key, title FROM project WHERE key = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	var project Project
	err = statement.QueryRow(key).Scan(&project.Key, &project.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &project, nil
}

func ListProjects() ([]Project, error) {
	db := InitDb()

	statement, err := db.Prepare("SELECT key, title FROM project ORDER BY key")
	if err != nil {
		return nil, err
	}

	var output []Project
	projects, err := statement.Query()
	if err != nil {
		return nil, err
	}
	for projects.Next() {
		var project Project
		err := projects.Scan(&project.Key, &project.Title)
		if err != nil {
			return nil, err
		}
		output = append(output, project)
	}

	return output, nil
}
