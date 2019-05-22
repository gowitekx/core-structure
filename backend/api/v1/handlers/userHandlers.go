package handlers

import (
	"net/http"
	"strconv"

	"github.com/infinity-framework/backend/api/v1/middleware"

	"github.com/gorilla/mux"
	"github.com/infinity-framework/backend/api/models"
	v1 "github.com/infinity-framework/backend/api/v1"
	"github.com/infinity-framework/backend/configs"

	"github.com/infinity-framework/backend/api"
	"github.com/infinity-framework/backend/api/v1/services"
)

//HTTPUserHandler struct
type HTTPUserHandler struct {
	UserService services.UserServices
}

//NewUserHTTPHandler Function
func NewUserHTTPHandler(userService services.UserServices, router api.Route) {
	handler := &HTTPUserHandler{UserService: userService}
	rm := middleware.RequestMiddleware{}

	router.Router.HandleFunc("/login", handler.UserLogin).Methods("POST", "OPTIONS")

	adminSubrouter := router.Router.PathPrefix("/admin/users").Subrouter()
	adminSubrouter.Use(rm.ValidateMiddleware) // User Admin User Middleware
	adminSubrouter.HandleFunc("/", handler.CreateUser).Methods("POST")
	adminSubrouter.HandleFunc("/{pageID}", handler.GetAllUsers).Methods("GET")
	adminSubrouter.HandleFunc("/id/:{email}", handler.GetUserByEmail).Methods("GET")
	adminSubrouter.HandleFunc("/:{email}", handler.UpdateUser).Methods("PUT")
	adminSubrouter.HandleFunc("/:{email}", handler.DeleteUser).Methods("DELETE")
	adminSubrouter.HandleFunc("/disable/:{email}", handler.DisableUser).Methods("PUT")

	userSubrouter := router.Router.PathPrefix("/user").Subrouter()
	userSubrouter.Use(rm.ValidateMiddleware)
	userSubrouter.HandleFunc("/:{email}", handler.GetUserByEmail).Methods("GET")
	userSubrouter.HandleFunc("/:{email}", handler.UpdateUser).Methods("PUT")
}

//UserLogin Method
func (httpUser HTTPUserHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	email, password, _ := r.BasicAuth()
	input := models.UserRequest{
		Email:    email,
		Password: password,
	}
	loginResponse, err1 := httpUser.UserService.UserLogin(r.Context(), &input)
	if err1 != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Read Body:", err1)
		v1.WriteErrorResponse(w, http.StatusUnauthorized, "Unable To Login!")
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   loginResponse.Token,
		Expires: loginResponse.ExpirationTime,
	})
	v1.WriteOKResponse(w, loginResponse)
}

//CreateUser Method
func (httpUser HTTPUserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	input := models.UserRequest{}
	defer r.Body.Close()
	err := v1.ReadInput(r.Body, &input)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Read Body:", err)
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Input For Creating User!")
		return
	}
	err1 := httpUser.UserService.CreateUser(r.Context(), &input)
	if err1 != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Read Body:", err)
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Unable To Create User!")
		return
	}
	v1.WriteOKResponse(w, "User Created Successfully!")
}

//GetAllUsers Method
func (httpUser HTTPUserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pageID := params["pageID"]
	page, err := strconv.Atoi(pageID)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Invalid Params:", err)
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid Input For Getting Users!")
		return
	}
	user, err := httpUser.UserService.GetAllUsers(r.Context(), page)
	if err != nil {
		v1.WriteErrorResponse(w, http.StatusNotFound, "No Users in the Database!")
		return
	}
	v1.WriteOKResponse(w, user)
}

//GetUserByEmail Method
func (httpUser HTTPUserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userEmail := params["email"]
	user, err := httpUser.UserService.GetUserByEmail(r.Context(), userEmail)
	if err != nil {
		v1.WriteErrorResponse(w, http.StatusNotFound, "User Not Found!")
		return
	}
	v1.WriteOKResponse(w, user)
}

//UpdateUser Method
func (httpUser HTTPUserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userEmail := params["email"]
	input := models.UserRequest{}
	defer r.Body.Close()
	err := v1.ReadInput(r.Body, &input)
	if err != nil {
		configs.Ld.Logger(r.Context(), configs.ERROR, "Read Body", err)
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Invaid Input!")
		return
	}
	err1 := httpUser.UserService.UpdateUser(r.Context(), userEmail, &input)
	if err1 != nil {
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Unable to Update User!")
		return
	}
	v1.WriteOKResponse(w, "User Updated Successfully!")
}

//DeleteUser Method
func (httpUser HTTPUserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userEmail := params["email"]
	err := httpUser.UserService.DeleteUser(r.Context(), userEmail)
	if err != nil {
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Unable to Delete User!")
		return
	}
	v1.WriteOKResponse(w, "User Deleted Successfully!")
}

//DisableUser Method
func (httpUser HTTPUserHandler) DisableUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userEmail := params["email"]
	err1 := httpUser.UserService.DisableUser(r.Context(), userEmail)
	if err1 != nil {
		v1.WriteErrorResponse(w, http.StatusBadRequest, "Unable to Disable User!")
		return
	}
	v1.WriteOKResponse(w, "User Disabled!")
}
