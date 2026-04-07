package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func NewCartDaoWithDb(db *gorm.DB) *CartDao {
	return &CartDao{db}
}

func (dao *CartDao) Create(cart *model.Cart) error {
	return dao.DB.Create(cart).Error
}

func (dao *CartDao) List(uid uint) (cart []*model.Cart, err error) {
	err = dao.DB.Where("user_id = ?", uid).Find(&cart).Error
	return
}

func (dao *CartDao) UpdateNum(id uint, num uint, uid uint) error {
	return dao.DB.Model(&model.Cart{}).Where("id = ? AND user_id = ?", id, uid).Update("num", num).Error
}

func (dao *CartDao) Delete(id uint, uid uint) error {
	return dao.DB.Where("id = ? AND user_id = ?", id, uid).Delete(&model.Cart{}).Error
}
