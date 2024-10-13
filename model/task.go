package model

type Task struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"Title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
