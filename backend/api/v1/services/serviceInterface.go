package services

import (
	"context"

	"github.com/infinity-framework/backend/api/models"
)

//UserServices Inteface
type UserServices interface {
	UserLogin(context.Context, *models.UserRequest) (LoginResponse, error)
	CreateUser(context.Context, *models.UserRequest) error
	GetAllUsers(context.Context, int) (AllUserResponse, error)
	GetUserByEmail(context.Context, string) (models.User, error)
	UpdateUser(context.Context, string, *models.UserRequest) error
	DeleteUser(context.Context, string) error
	DisableUser(context.Context, string) error
}

//CoursesServices Intefaces
type CoursesServices interface {
	CreateCourse(context.Context, *models.Courses) error
	UpdateCourse(context.Context, *models.Courses, int) error
	DeleteCourse(context.Context, int) error
	GetAllCourses(context.Context, int) (AllCoursesResponse, error)
	GetCourseByID(context.Context, int) (models.Courses, error)
}
