package models

type ListProjectsRequest struct {
	TeamID int `json:"team_id"`
}

type ListProjectsResponse struct {
	ID          uint   `json:"id"`
	ProjectName string `json:"project_name"`
	Icon        string `json:"icon"`
	ProjectType string `json:"project_type"`
}

type CreateProjectRequest struct {
	ProjectName string `json:"project_name"`
	Icon        string `json:"icon"`
	ProjectType string `json:"project_type"`
	TeamID      int    `json:"team_id"`
}

type DeleteProjectsRequest struct {
	ID int `json:"id"`
}

type UpdateProjectRequest struct {
	ID          uint   `json:"id"`
	ProjectName string `json:"project_name"`
	Icon        string `json:"icon"`
}

type AddFavoriteProjectRequest struct {
	ProjectID int `json:"project_id"`
	UserID    int `json:"user_id"`
}
