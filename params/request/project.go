package request

type ProjectCreateRequest struct {
	ProjectName string `json:"projectName"`
	ProjectDesc string `json:"projectDesc"`
	CreatedBy   string `json:"createdBy"`
}

type ProejctListRequest struct {
	PageSize int `form:"pageSize"`
	PageNum  int `form:"pageNum"`
}
