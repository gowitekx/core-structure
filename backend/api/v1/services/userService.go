package services

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/infinity-framework/backend/api/models"
	"github.com/infinity-framework/backend/api/v1/repository"
	"github.com/infinity-framework/backend/configs"
)

type AllUserResponse struct {
	Users []models.User
}

//LoginResponse Struct
type LoginResponse struct {
	Token          string    `json:"token,omitempty"`
	UserType       string    `json:"userType,omitempty"`
	ExpirationTime time.Time `json:"-"`
	Error          string    `json:"error,omitempty"`
}

//Claims Struct
type Claims struct {
	Email    string `json:"email"`
	UserType string `json:"userType"`
	jwt.StandardClaims
}

//UserService to manage users persistence
type UserService struct {
	userRepository repository.UserRepository
}

//NewUserService func
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

//UserLogin Func
func (u UserService) UserLogin(ctx context.Context, input *models.UserRequest) (LoginResponse, error) {
	loginResponse := LoginResponse{}
	userData := models.User{}
	userData.Email = input.Email
	userData.Password = input.Password

	userDetails, err := u.userRepository.UserLogin(ctx, &userData)
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "SERVICE:Login: ", err)
		return loginResponse, err
	}

	errp := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(input.Password))
	if errp != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "SERVICE:Password Comapare: ", errp)
		return loginResponse, errp
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		Email:    userDetails.Email,
		UserType: userDetails.UserType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, error := token.SignedString([]byte(configs.Config.JWTSecretKey))
	if error != nil {
		return loginResponse, error
	}
	loginResponse.Token = tokenString
	loginResponse.UserType = userDetails.UserType
	return loginResponse, nil
}

//CreateUser Func
func (u UserService) CreateUser(ctx context.Context, input *models.UserRequest) error {
	userData := models.User{}
	pass, errp := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if errp != nil {
		return errp
	}
	userData.Email = input.Email
	userData.Name = input.Name
	userData.Password = string(pass)
	userData.Designation = input.Designation
	userData.EmpID = input.EmpID
	userData.UserType = input.UserType
	userData.UserStatus = input.UserStatus
	userData.CreatedAt = time.Now()
	userData.UpdatedAt = time.Now()
	err := u.userRepository.CreateUser(ctx, &userData)
	if err != nil {
		return err
	}
	return nil
}

//GetAllUsers Func
func (u UserService) GetAllUsers(ctx context.Context, page int) (AllUserResponse, error) {
	var allUserResponse AllUserResponse
	page = (page - 1) * 10
	userData, err := u.userRepository.GetAllUsers(ctx, page)
	if err != nil {
		return allUserResponse, err
	}
	user := models.User{}
	for _, value := range userData {
		user.ID = value.ID
		user.Email = value.Email
		user.Name = value.Name
		user.Designation = value.Designation
		user.EmpID = value.EmpID
		user.UserType = value.UserType
		user.UserStatus = value.UserStatus
		user.CreatedAt = value.CreatedAt
		allUserResponse.Users = append(allUserResponse.Users, user)
	}
	return allUserResponse, nil
}

//GetUserByEmail Func
func (u UserService) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	user := models.User{}
	userData, err := u.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return user, err
	}
	user.ID = userData.ID
	user.Email = userData.Email
	user.Name = userData.Name
	user.Designation = userData.Designation
	user.EmpID = userData.EmpID
	user.UserType = userData.UserType
	user.UserStatus = userData.UserStatus
	user.CreatedAt = userData.CreatedAt
	return user, nil
}

//UpdateUser Func
func (u UserService) UpdateUser(ctx context.Context, email string, input *models.UserRequest) error {
	userData := models.User{}
	userData.Email = input.Email
	userData.Name = input.Name
	userData.Password = input.Password
	userData.Designation = input.Designation
	userData.EmpID = input.EmpID
	userData.UserType = input.UserType
	userData.UserStatus = input.UserStatus
	userData.UpdatedAt = time.Now()
	err := u.userRepository.UpdateUser(ctx, email, &userData)
	if err != nil {
		return err
	}
	return nil
}

//DeleteUser Func
func (u UserService) DeleteUser(ctx context.Context, email string) error {
	err := u.userRepository.DeleteUser(ctx, email)
	if err != nil {
		return err
	}
	return nil
}

//DisableUser Func
func (u UserService) DisableUser(ctx context.Context, email string) error {
	err := u.userRepository.DisableUser(ctx, email)
	if err != nil {
		return err
	}
	return nil
}
