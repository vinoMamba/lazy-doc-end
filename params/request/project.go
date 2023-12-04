package request

type ProjectCreateRequest struct {
	ProjectName string `json:"projectName"`
	ProjectDesc string `json:"projectDesc"`
	IsPublic    int    `json:"isPublic"`
}

type ProejctListRequest struct {
	PageSize int `form:"pageSize"`
	PageNum  int `form:"pageNum"`
}
