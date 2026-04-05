package service

import (
	"context"
	"fmt"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/pkg/util"
	"gin_mall/serializer"
	"strconv"
)

type FavoriteService struct {
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	ProductId  uint `json:"product_id" form:"product_id"`
	model.BasePage
}

func (service *FavoriteService) List(ctx context.Context, uid uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	code := e.Success
	favorite, err := favoriteDao.List(uid)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorite), uint(len(favorite)))
}

func (service *FavoriteService) Creat(ctx context.Context, uid uint) serializer.Response {
	fmt.Println(service.BossId, uid)
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	exist, _ := favoriteDao.ExistOrNot(service.ProductId, uid)
	if exist {
		code := e.ErrorExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  e.GetMsg(code),
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByID(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.FindUserByID(service.BossId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	favorite := &model.Favorite{
		UserId:    user.ID,
		BossId:    service.BossId,
		ProductId: service.ProductId,
		User:      *user,
		Boss:      *boss,
		Product:   *product,
	}
	err = favoriteDao.Create(favorite)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildFavorite(favorite, product, boss),
	}
}

func (service *FavoriteService) Delete(ctx context.Context, uid uint, fid string) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	id, _ := strconv.ParseUint(fid, 10, 0)
	err := favoriteDao.Delete(uid, uint(id))
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
