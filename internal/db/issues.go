package db

type Issue struct {
	ID          uint
	ProjectKey  string
	Title       string
	Description string
}

func ListIssuesForProject(key string) ([]Issue, error) {
	db := InitDb()

	statement, err := db.Prepare("SELECT id, project_key, title, description FROM issue WHERE project_key = ? ORDER BY id")
	if err != nil {
		return nil, err
	}

	var output []Issue
	issues, err := statement.Query(key)
	if err != nil {
		return nil, err
	}
	for issues.Next() {
		var issue Issue
		err := issues.Scan(&issue.ID, &issue.ProjectKey, &issue.Title, &issue.Description)
		if err != nil {
			return nil, err
		}
		output = append(output, issue)
	}

	return output, nil
}
