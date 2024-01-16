package request

type CreateDirRequest struct {
	DirName  string `json:"dirName" valid:"required"`
	ParentId int64  `json:"parentId"`
	Status   int    `json:"status"`
}
