package request

type CreateDirRequest struct {
	DirName  string `json:"dirName" valid:"required"`
	ParentId string `json:"parentId"`
	IsPublic int    `json:"isPublic"`
}

type UpdateDirRequest struct {
	DirName  string `json:"dirName" valid:"required"`
	ParentId int64  `json:"parentId"`
	IsPublic int    `json:"isPublic"`
}

type DirListRequest struct {
	Id       string `json:"id"`
	ParentId string `json:"parentId"`
	DirName  string `json:"dirName"`
	IsPublic int    `json:"isPublic"`
}
