package models

type ListTeamRequest struct {
	UserID int `json:"user_id"`
}

type ListTeamResponse struct {
	ID       uint   `json:"id"`
	TeamName string `json:"team_name"`
	Icon     string `json:"icon"`
}
