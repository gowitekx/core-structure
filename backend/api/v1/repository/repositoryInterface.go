package repository

import (
	"context"

	"github.com/infinity-framework/backend/api/models"
)

//UserRepository Interface
type UserRepository interface {
	UserLogin(context.Context, *models.User) (*models.User, error)
	CreateUser(context.Context, *models.User) error
	GetAllUsers(context.Context, int) ([]models.User, error)
	GetUserByEmail(context.Context, string) (models.User, error)
	UpdateUser(context.Context, string, *models.User) error
	DeleteUser(context.Context, string) error
	DisableUser(context.Context, string) error
}

//CoursesRepository Inteface
type CoursesRepository interface {
	CreateCourse(context.Context, *models.Courses) error
	UpdateCourse(context.Context, *models.Courses, int) error
	DeleteCourse(context.Context, int) error
	GetAllCourses(context.Context, int) ([]models.Courses, error)
	GetCourseByID(context.Context, int) (models.Courses, error)
}
