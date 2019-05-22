package services

import (
	"context"
	"time"

	"github.com/infinity-framework/backend/api/models"
	"github.com/infinity-framework/backend/api/v1/repository"
)

//AllCoursesResponse Struct
type AllCoursesResponse struct {
	Courses []models.Courses
}

//CoursesService Service
type CoursesService struct {
	coursesRepository repository.CoursesRepository
}

//NewCoursesService func
func NewCoursesService(coursesRepo repository.CoursesRepository) *CoursesService {
	return &CoursesService{coursesRepository: coursesRepo}
}

//CreateCourse Func
func (c CoursesService) CreateCourse(ctx context.Context, input *models.Courses) error {
	courseData := models.Courses{}
	courseData.Name = input.Name
	courseData.Description = input.Description
	courseData.CreatedAt = time.Now()
	courseData.UpdatedAt = time.Now()
	err := c.coursesRepository.CreateCourse(ctx, &courseData)
	if err != nil {
		return err
	}
	return nil
}

//UpdateCourse Func
func (c CoursesService) UpdateCourse(ctx context.Context, input *models.Courses, id int) error {
	courseData := models.Courses{}
	courseData.Name = input.Name
	courseData.Description = input.Description
	courseData.UpdatedAt = time.Now()
	err := c.coursesRepository.UpdateCourse(ctx, &courseData, id)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCourse Func
func (c CoursesService) DeleteCourse(ctx context.Context, id int) error {
	err := c.coursesRepository.DeleteCourse(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

//GetAllCourses Func
func (c CoursesService) GetAllCourses(ctx context.Context, page int) (AllCoursesResponse, error) {
	var allCoursesResponse AllCoursesResponse
	page = (page - 1) * 10
	coursesData, err := c.coursesRepository.GetAllCourses(ctx, page)
	if err != nil {
		return allCoursesResponse, err
	}
	courses := models.Courses{}
	for _, value := range coursesData {
		courses.ID = value.ID
		courses.Name = value.Name
		courses.Description = value.Description
		allCoursesResponse.Courses = append(allCoursesResponse.Courses, courses)
	}
	return allCoursesResponse, nil
}

//GetCourseByID Func
func (c CoursesService) GetCourseByID(ctx context.Context, id int) (models.Courses, error) {
	courseData := models.Courses{}
	course, err := c.coursesRepository.GetCourseByID(ctx, id)
	if err != nil {
		return course, err
	}
	courseData.ID = course.ID
	courseData.Name = course.Name
	courseData.Description = course.Description
	return course, nil
}
