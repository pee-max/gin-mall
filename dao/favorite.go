package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoWithDb(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

func (dao *FavoriteDao) List(uid uint) (favorites []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id = ?", uid).Find(&favorites).Error
	return
}

func (dao *FavoriteDao) ExistOrNot(pid uint, uid uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("user_id = ? AND product_id = ?", uid, pid).Count(&count).Error
	if err != nil {
		return false, nil
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (dao *FavoriteDao) Create(favorite *model.Favorite) (err error) {
	err = dao.DB.Model(&model.Favorite{}).Create(favorite).Error
	return
}

func (dao *FavoriteDao) Delete(uid uint, fid uint) (err error) {
	err = dao.DB.Model(&model.Favorite{}).
		Where("id = ? AND user_id = ?", fid, uid).Delete(&model.Favorite{}).Error
	return
}
