package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoWithDb(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

func (dao *AddressDao) Create(address *model.Address) error {
	return dao.DB.Create(address).Error
}
func (dao *AddressDao) FindById(id uint, uid uint) (address *model.Address, err error) {
	err = dao.DB.Where("id = ? AND user_id = ?", id, uid).First(&address).Error
	return
}

func (dao *AddressDao) List(uid uint) (address []*model.Address, err error) {
	err = dao.DB.Where("user_id = ?", uid).Find(&address).Error
	return
}

func (dao *AddressDao) Update(address *model.Address, id, uid uint) error {
	return dao.DB.Where("id = ? AND user_id = ?", id, uid).Updates(address).Error
}

func (dao *AddressDao) Delete(id uint, uid uint) error {
	return dao.DB.Where("id = ? AND user_id = ?", id, uid).Delete(&model.Address{}).Error
}
