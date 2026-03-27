package dao

import (
	"context"
	"errors"
	"fmt"
	"gin_mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoWithDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).First(&user).Error
	//fmt.Println(err, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		fmt.Printf("database error: %v \n", err)
		return nil, false, err
	}
	return user, true, nil
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Create(user).Error
}

func (dao *UserDao) FindUserByID(uId uint) (user *model.User, err error) {
	user = &model.User{}
	err = dao.DB.Where("id=?", uId).First(user).Error
	return user, err
}

func (dao *UserDao) UpdateUserByID(uId uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uId).Updates(&user).Error
}
