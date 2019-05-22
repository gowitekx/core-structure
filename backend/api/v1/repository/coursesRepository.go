package repository

import (
	"context"

	"github.com/infinity-framework/backend/api/models"
	"github.com/infinity-framework/backend/configs"
	"github.com/infinity-framework/backend/database/connection"
	"github.com/jinzhu/gorm"
)

//CoursesConnRepository Struct
type CoursesConnRepository struct {
	ConnectionService connection.ConnectionInterface
	DB                *gorm.DB
}

//NewCoursesRepository Func
func NewCoursesRepository(connectionService connection.ConnectionInterface, db *gorm.DB) *CoursesConnRepository {
	return &CoursesConnRepository{connectionService, db}
}

//CreateCourse Func
func (c CoursesConnRepository) CreateCourse(ctx context.Context, course *models.Courses) error {
	err := c.DB.Table("courses").Create(&course).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:Cannot Create Course: ", err)
		return err
	}
	return nil
}

//UpdateCourse Func
func (c CoursesConnRepository) UpdateCourse(ctx context.Context, course *models.Courses, id int) error {
	err := c.DB.Table("courses").Where("id=?", id).Update(&course).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:Cannot Create Course: ", err)
		return err
	}
	return nil
}

//DeleteCourse Func
func (c CoursesConnRepository) DeleteCourse(ctx context.Context, id int) error {
	course := models.Courses{}
	err := c.DB.Table("courses").Where("id=?", id).Delete(&course).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:Cannot Create Course: ", err)
		return err
	}
	return nil
}

//GetAllCourses Func
func (c CoursesConnRepository) GetAllCourses(ctx context.Context, page int) ([]models.Courses, error) {
	courses := []models.Courses{}
	err := c.DB.Table("courses").Limit(10).Offset(page).Find(&courses).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:No Courses Found In DB: ", err)
		return courses, err
	}
	return courses, nil
}

//GetCourseByID Func
func (c CoursesConnRepository) GetCourseByID(ctx context.Context, id int) (models.Courses, error) {
	course := models.Courses{}
	err := c.DB.Table("courses").Where("id=?", id).First(&course).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:Course Not Found In DB: ", err)
		return course, err
	}
	return course, nil
}
