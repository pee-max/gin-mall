package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func NewOrderDaoWithDb(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

func (dao *OrderDao) Create(order *model.Order) error {
	return dao.DB.Create(order).Error
}
func (dao *OrderDao) FindById(uid uint, id uint) (order *model.Order, err error) {
	err = dao.DB.Where("id = ? AND user_id = ?", id, uid).First(&order).Error
	return
}

func (dao *OrderDao) List(uid uint) (order []*model.Order, err error) {
	err = dao.DB.Where("user_id = ?", uid).Find(&order).Error
	return
}

func (dao *OrderDao) ListWithCondition(condition map[string]interface{}, page model.BasePage) (orders []*model.Order, err error) {
	err = dao.DB.Where(condition).Offset((page.PageSize) * (page.PageNum - 1)).Limit(page.PageSize).Find(&orders).Error
	return
}

func (dao *OrderDao) Delete(id uint, uid uint) error {
	return dao.DB.Where("id = ? AND user_id = ?", id, uid).Delete(&model.Order{}).Error
}
