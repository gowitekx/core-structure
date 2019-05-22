package main

import (
	"log"
	"net/http"

	handler "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/infinity-framework/backend/api"
	"github.com/infinity-framework/backend/api/v1/handlers"
	"github.com/infinity-framework/backend/api/v1/middleware"
	"github.com/infinity-framework/backend/api/v1/repository"
	"github.com/infinity-framework/backend/api/v1/services"
	"github.com/infinity-framework/backend/configs"
	"github.com/infinity-framework/backend/database"
	"github.com/infinity-framework/backend/database/connection"
)

//initialize the code
func init() {
	//Set Env variable for config file
	configs.Config.Read("production")
}

// We Manage / Set Environment in config.toml file
const (
	DEVELOPMENT = iota
	TEST
	STAGE
	PRODUCTION
)

func main() {
	log.Println("Server Running On Port:", configs.Config.Port)
	database.DBMigrate() // DB Migration
	router := api.Route{}
	router.Router = mux.NewRouter()
	rm := middleware.RequestMiddleware{}
	router.Router.Use(rm.RequestIDGenerator)

	//Create Database Connection
	connectionService := connection.NewDatabaseConnection()
	db := connectionService.DBConnect()

	//User
	userRepository := repository.NewUserRepository(connectionService, db)
	userService := services.NewUserService(userRepository)
	handlers.NewUserHTTPHandler(userService, router)

	//Courses
	coursesRepository := repository.NewCoursesRepository(connectionService, db)
	courseService := services.NewCoursesService(coursesRepository)
	handlers.NewCoursesHTTPHandler(courseService, router)

	log.Fatal(http.ListenAndServe(":"+configs.Config.Port, handler.CORS(handler.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handler.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handler.AllowedOrigins([]string{"*"}))(router.Router)))
}
