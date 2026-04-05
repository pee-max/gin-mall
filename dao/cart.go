package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type CartDao struct {
	db *gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func NewCartDaoWithDb(db *gorm.DB) *CartDao {
	return &CartDao{db}
}

func (dao *CartDao) Create(cart *model.Cart) error {
	return dao.db.Create(cart).Error
}

func (dao *CartDao) List(uid uint) (cart []*model.Cart, err error) {
	err = dao.db.Where("user_id = ?", uid).Find(&cart).Error
	return
}

func (dao *CartDao) UpdateNum(id uint, num uint, uid uint) error {
	return dao.db.Model(&model.Cart{}).Where("id = ? AND user_id = ?", id, uid).Update("num", num).Error
}

func (dao *CartDao) Delete(id uint, uid uint) error {
	return dao.db.Where("id = ? AND user_id = ?", id, uid).Delete(&model.Cart{}).Error
}
