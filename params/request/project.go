package request

type ProjectCreateRequest struct {
	ProjectName string `json:"projectName"`
	ProjectDesc string `json:"projectDesc"`
	IsPublic    string `json:"isPublic"`
}

type ProejctListRequest struct {
	PageSize int `form:"pageSize"`
	PageNum  int `form:"pageNum"`
}

type ProjectUpdateRequest struct {
	Id          int    `json:"id"`
	ProjectName string `json:"projectName"`
	ProjectDesc string `json:"projectDesc"`
	IsPublic    string `json:"isPublic"`
}
