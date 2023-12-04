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
	ProjectName string
	ProejctDesc string `gorm:"column:project_description"`
	IsPublic    int    `gorm:"column:is_public"`
	CreatedBy   int64  `gorm:"column:created_by"`
	CommonTimestampsFields
}
