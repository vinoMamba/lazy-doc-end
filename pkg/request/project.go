package request

type CreateProjectRequest struct {
	ProjectName string `json:"projectName" valid:"required"`
	ProjectDesc string `json:"projectDesc"`
	Status      int    `json:"status"`
}
