package request

type ProjectCreateRequest struct {
	ProjectName string `json:"projectName"`
	ProjectDesc string `json:"projectDesc"`
	CreatedBy   string `json:"createdBy"`
}
