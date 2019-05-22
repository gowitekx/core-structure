package repository

import (
	"context"

	"github.com/infinity-framework/backend/api/models"
	"github.com/infinity-framework/backend/configs"
	"github.com/infinity-framework/backend/database/connection"
	"github.com/jinzhu/gorm"
)

//UserConnRepository Struct
type UserConnRepository struct {
	ConnectionService connection.ConnectionInterface
	DB                *gorm.DB
}

//NewUserRepository Func
func NewUserRepository(connectionService connection.ConnectionInterface, db *gorm.DB) *UserConnRepository {
	return &UserConnRepository{connectionService, db}
}

//UserLogin Func
func (u UserConnRepository) UserLogin(ctx context.Context, user *models.User) (*models.User, error) {
	err := u.DB.Table("users").Where("email=?", user.Email).First(&user).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:Cannot Login: ", err)
		return user, err
	}
	return user, nil
}

//CreateUser Func
func (u UserConnRepository) CreateUser(ctx context.Context, user *models.User) error {

	err := u.DB.Table("users").Create(&user).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:Cannot Create User: ", err)
		return err
	}
	return nil
}

//GetAllUsers Func
func (u UserConnRepository) GetAllUsers(ctx context.Context, page int) ([]models.User, error) {
	users := []models.User{}
	err := u.DB.Table("users").Limit(10).Offset(page).Find(&users).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:No Users Found In DB: ", err)
		return users, err
	}
	return users, nil
}

//GetUserByEmail Func
func (u UserConnRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	user := models.User{}
	err := u.DB.Table("users").Where("email=?", email).First(&user).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:User Not Found In DB: ", err)
		return user, err
	}
	return user, nil
}

//UpdateUser Func
func (u UserConnRepository) UpdateUser(ctx context.Context, email string, user *models.User) error {
	err := u.DB.Table("users").Where("email=?", email).Update(&user).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:User Not Found In DB: ", err)
		return err
	}
	return nil
}

//DeleteUser Func
func (u UserConnRepository) DeleteUser(ctx context.Context, email string) error {
	user := models.User{}
	err := u.DB.Table("users").Where("email=?", email).Delete(&user).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:User Not Found In DB: ", err)
		return err
	}
	return nil
}

//DisableUser Func
func (u UserConnRepository) DisableUser(ctx context.Context, email string) error {
	err := u.DB.Table("users").Where("email=?", email).Update("user_status", false).Error
	if err != nil {
		configs.Ld.Logger(ctx, configs.ERROR, "REPO:User Not Found In DB: ", err)
		return err
	}
	return nil
}
