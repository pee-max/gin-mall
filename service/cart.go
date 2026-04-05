package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
	"strconv"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       uint `json:"num" form:"num"`
}

func (service *CartService) Create(ctx context.Context, uid uint) serializer.Response {
	var cart *model.Cart
	productDao := dao.NewProductDao(ctx)
	code := e.Success
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	cartDao := dao.NewCartDao(ctx)
	cart = &model.Cart{
		UserId:    uid,
		ProductId: product.ID,
		BossId:    product.BossId,
		Num:       service.Num,
	}
	err = cartDao.Create(cart)
	if err != nil {
		code = e.ErrorFailCreat
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.FindUserByID(cart.BossId)
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
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

func (service *CartService) List(ctx context.Context, uid uint) serializer.Response {
	var carts []*model.Cart
	cartDao := dao.NewCartDao(ctx)
	code := e.Success
	carts, err := cartDao.List(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarts(ctx, carts), uint(len(carts)))
}

func (service *CartService) Update(ctx context.Context, uid uint, id string) serializer.Response {
	cartDao := dao.NewCartDao(ctx)
	code := e.Success
	cid, _ := strconv.Atoi(id)
	err := cartDao.UpdateNum(uint(cid), service.Num, uid)
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
	}
}

func (service *CartService) Delete(ctx context.Context, uid uint, id string) serializer.Response {
	cartDao := dao.NewCartDao(ctx)
	code := e.Success
	cid, _ := strconv.Atoi(id)
	err := cartDao.Delete(uint(cid), uint(uid))
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
	}
}
