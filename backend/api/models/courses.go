package models

import "time"

//Courses Model
type Courses struct {
	ID          int        `gorm:"id" json:"id"`
	Name        string     `gorm:"name" json:"courseName"`
	Description string     `gorm:"description" json:"description"`
	CreatedAt   time.Time  `gorm:"created_at" json:"-"`
	UpdatedAt   time.Time  `gorm:"updated_at" json:"-"`
	DeletedAt   *time.Time `gorm:"deleted_at" json:"-"`
}
