package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/pkg/e"
	"gin_mall/pkg/util"
	"gin_mall/serializer"
)

type CategoryService struct {
}

func (service *CategoryService) List(ctx context.Context) serializer.Response {
	CategoryDao := dao.NewCategoryDao(ctx)
	code := e.Success
	Categories, err := CategoryDao.ListCategory()
	if err != nil {
		util.LogrusObj.Infoln("[err] ", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategories(Categories), uint(len(Categories)))
}
