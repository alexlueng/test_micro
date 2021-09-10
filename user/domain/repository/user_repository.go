package repository

import (
	"micro/user/domain/model"

	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	InitTable() error
	FindUserByName(string) (*model.User, error)
	FindUserByID(int64) (*model.User, error)
	CreateUser(*model.User) (int64, error)
	DeleteUserByID(int64) error
	UpdateUser(*model.User) error
	FindAll() ([]model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		mysqlDB: db,
	}
}

type UserRepository struct {
	mysqlDB *gorm.DB
}

func (u *UserRepository) InitTable() error {
	return u.mysqlDB.CreateTable(&model.User{}).Error
}

func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDB.Where("user_name = ?", name).Find(user).Error
}

func (u *UserRepository) FindUserByID(id int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDB.Find(user, id).Error
}

func (u *UserRepository) CreateUser(user *model.User) (userID int64, err error) {
	return user.ID, u.mysqlDB.Create(user).Error
}

func (u *UserRepository) DeleteUserByID(id int64) error {
	return u.mysqlDB.Where("id = ?", id).Delete(&model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDB.Model(user).Update(&user).Error
}

func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDB.Find(&userAll).Error
}
