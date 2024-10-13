package controller

type taskRequest struct {
	Title       string `gorm:"not null" json:"Title"`
	Description string `json:"description"`
}
