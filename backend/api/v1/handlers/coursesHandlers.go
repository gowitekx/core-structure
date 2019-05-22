package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/infinity-framework/backend/api/models"
	v1 "github.com/infinity-framework/backend/api/v1"
	"github.com/infinity-framework/backend/configs"

	"github.com/infinity-framework/backend/api"
	"github.com/infinity-framework/backend/api/v1/middleware"
	"github.com/infinity-framework/backend/api/v1/services"
)

//HTTPCoursesHandler Struct
type HTTPCoursesHandler struct {
	CoursesService services.CoursesServices
}

//NewCoursesHTTPHandler Function
func NewCoursesHTTPHandler(coursesService services.CoursesServices, router api.Route) {
	handler := &HTTPCoursesHandler{CoursesService: coursesService}

	rm := middleware.RequestMiddleware{}

	courseSubrouter := router.Router.PathPrefix("/admin/courses").Subrouter()
	courseSubrouter.Use(rm.ValidateMiddleware)
	courseSubrouter.HandleFunc("/", handler.CreateCourse).Methods("POST")
	courseSubrouter.HandleFunc("/:id={id}", handler.UpdateCourse).Methods("PUT")
	courseSubrouter.HandleFunc("/:id={id}", handler.DeleteCourse).Methods("DELETE")
	courseSubrouter.HandleFunc("/all/{pageID}", handler.GetAllCourses).Methods("GET")
	courseSubrouter.HandleFunc("/:id={id}", handler.GetCourseByID).Methods("GET")

	courseUserSubrouter := router.Router.PathPrefix("/user/courses").Subrouter()
	courseUserSubrouter.Use(rm.ValidateMiddleware)
	courseUserSubrouter.HandleFunc("/:id={id}", handler.GetCourseByID).Methods("GET")
}

//CreateCourse Method
func (httpCourses HTTPCoursesHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	input := models.Courses{}
	defer r.Body.Close()
	err := v1.ReadInput(r.Body, &input)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Read Body:", err)
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Input!")
		return
	}
	err1 := httpCourses.CoursesService.CreateCourse(r.Context(), &input)
	if err1 != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Read Body:", err1)
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Unable To Create Course!")
		return
	}
	v1.WriteOKResponse(w, "Course Created Successfully!")
}

//UpdateCourse Method
func (httpCourses HTTPCoursesHandler) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	courseID := params["id"]
	cid, err := strconv.Atoi(courseID)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Invalid Params:", err)
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Input!")
		return
	}
	input := models.Courses{}
	defer r.Body.Close()
	err1 := v1.ReadInput(r.Body, &input)
	if err1 != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Read Body:", err1)
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Input!")
		return
	}
	err2 := httpCourses.CoursesService.UpdateCourse(r.Context(), &input, cid)
	if err2 != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Read Body:", err2)
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Unable To Update Course!")
		return
	}
	v1.WriteOKResponse(w, "Course Updated Successfully!")
}

//DeleteCourse Method
func (httpCourses HTTPCoursesHandler) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	courseID := params["id"]
	cid, err := strconv.Atoi(courseID)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Invalid Params:", err)
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Input!")
		return
	}
	err1 := httpCourses.CoursesService.DeleteCourse(r.Context(), cid)
	if err1 != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Read Body:", err1)
		v1.WriteErrorResponse(w, http.StatusNotFound, "Course Not Found!")
		return
	}
	v1.WriteOKResponse(w, "Course Deleted Successfully!")
}

//GetAllCourses Method
func (httpCourses HTTPCoursesHandler) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pageID := params["pageID"]
	page, err := strconv.Atoi(pageID)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Invalid Params:", err)
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Input!")
		return
	}
	courses, err1 := httpCourses.CoursesService.GetAllCourses(r.Context(), page)
	if err1 != nil {
		v1.WriteErrorResponse(w, http.StatusNotFound, "No Courses in the Database!")
		return
	}
	v1.WriteOKResponse(w, courses)
}

//GetCourseByID Method
func (httpCourses HTTPCoursesHandler) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	courseID := params["id"]
	cid, err := strconv.Atoi(courseID)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Invalid Params:", err)
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Input!")
		return
	}
	course, err1 := httpCourses.CoursesService.GetCourseByID(r.Context(), cid)
	if err1 != nil {
		v1.WriteErrorResponse(w, http.StatusNotFound, "Course Not Found!")
		return
	}
	v1.WriteOKResponse(w, course)
}
