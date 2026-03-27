package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoWithDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

func (dao *NoticeDao) FindNoticeByID(uId uint) (Notice *model.Notice, err error) {
	err = dao.DB.Where("id=?", uId).First(&Notice).Error
	return
}
