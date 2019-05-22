package models

import "time"

//User Model
type User struct {
	ID          int        `gorm:"id" json:"id"`
	Email       string     `gorm:"email" json:"email"`
	Name        string     `gorm:"name" json:"name"`
	Password    string     `gorm:"password" json:"-"`
	Designation string     `gorm:"designation" json:"designation"`
	EmpID       string     `gorm:"emp_id" json:"empID"`
	UserType    string     `gorm:"user_type" json:"userType"`
	UserStatus  bool       `gorm:"user_status" json:"userStatus"`
	CreatedAt   time.Time  `gorm:"created_at" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"updated_at" json:"-"`
	DeletedAt   *time.Time `gorm:"deleted_at" json:"-"`
}

//UserRequest Type
type UserRequest struct {
	Email       string `gorm:"email" json:"email"`
	Name        string `gorm:"name" json:"name"`
	Password    string `gorm:"password" json:"password"`
	Designation string `gorm:"designation" json:"designation"`
	EmpID       string `gorm:"emp_id" json:"empID"`
	UserType    string `gorm:"user_type" json:"userType"`
	UserStatus  bool   `gorm:"user_status" json:"userStatus"`
}
