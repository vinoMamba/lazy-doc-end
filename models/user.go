package models

type User struct {
	BaseModel
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CommonTimestampsFields
}

type Project struct {
	BaseModel
	ProjectName string `gorm:"column:project_name" json:"projectName"`
	ProejctDesc string `gorm:"column:project_description" json:"projectDesc"`
	IsPublic    string `gorm:"column:is_public" json:"isPublic"`
	CreatedBy   int64  `gorm:"column:created_by" json:"-"`
	CommonTimestampsFields
}
