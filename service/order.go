package service

import (
	"context"
	"fmt"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
	"math/rand"
	"strconv"
	"time"
)

type OrderService struct {
	ProductId uint    `json:"product_id" form:"product_id"`
	Num       uint    `json:"num" form:"num"`
	AddressId uint    `json:"address_id" form:"address_id"`
	Money     float64 `json:"money" form:"money"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	UserId    uint    `json:"user_id" form:"user_id"`
	OrderNum  uint    `json:"order_num" form:"order_num"`
	Type      int     `json:"type" form:"type"`
	model.BasePage
}

func (service *OrderService) Create(ctx context.Context, uid uint) serializer.Response {
	var order *model.Order
	orderDao := dao.NewOrderDao(ctx)
	code := e.Success
	order = &model.Order{
		UserId:    uid,
		ProductId: service.ProductId,
		Num:       service.Num,
		Type:      1, //unpaid
	}
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.FindById(service.AddressId, uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.AddressId = address.ID

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
	order.BossId = product.BossId
	discountPrice, _ := strconv.ParseFloat(product.DiscountPrice, 64)
	order.Money = discountPrice * float64(order.Num)

	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000))
	productNum := strconv.Itoa(int(service.ProductId))
	userNum := strconv.Itoa(int(order.UserId))
	number = number + userNum + productNum
	order.OrderNum, _ = strconv.ParseUint(number, 10, 0)

	err = orderDao.Create(order)
	if err != nil {
		code = e.ErrorFailCreat
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

func (service *OrderService) Get(ctx context.Context, uid uint, id string) serializer.Response {
	var order *model.Order
	orderDao := dao.NewOrderDao(ctx)
	code := e.Success
	oid, _ := strconv.Atoi(id)
	order, err := orderDao.FindById(uid, uint(oid))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.FindById(order.AddressId, uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductId)
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
		Data:   serializer.BuildOrder(order, address, product),
	}
}

func (service *OrderService) List(ctx context.Context, uid uint) serializer.Response {
	var orders []*model.Order
	orderDao := dao.NewOrderDao(ctx)
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	condition := make(map[string]interface{})
	if service.Type != 0 {
		condition["type"] = service.Type
	}
	condition["user_id"] = uid

	orders, err := orderDao.ListWithCondition(condition, service.BasePage)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildOrders(ctx, orders), uint(len(orders)))
}

func (service *OrderService) Delete(ctx context.Context, uid uint, id string) serializer.Response {
	orderDao := dao.NewOrderDao(ctx)
	code := e.Success
	oid, _ := strconv.Atoi(id)
	err := orderDao.Delete(uint(oid), uid)
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
