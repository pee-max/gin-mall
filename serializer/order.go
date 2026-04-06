package serializer

import (
	"context"
	"gin_mall/conf"
	"gin_mall/dao"
	"gin_mall/model"
)

type Order struct {
	Id            uint    `json:"id"`
	OrderNum      uint64  `json:"order_num"`
	CreatedAt     int64   `json:"created_at"`
	UpdatedAt     int64   `json:"updated_at"`
	UserId        uint    `json:"user_id"`
	ProductId     uint    `json:"product_id"`
	BossId        uint    `json:"boss_id"`
	Num           uint    `json:"num"`
	AddressName   string  `json:"address_name"`
	AddressPhone  string  `json:"address_phone"`
	Address       string  `json:"address"`
	Type          uint    `json:"type"`
	ProductName   string  `json:"product_name"`
	ImgPath       string  `json:"img_path"`
	DiscountPrice float64 `json:"discount_price"`
}

func BuildOrder(order *model.Order, address *model.Address, product *model.Product) *Order {
	return &Order{
		Id:            order.ID,
		OrderNum:      order.OrderNum,
		CreatedAt:     order.CreatedAt.Unix(),
		UpdatedAt:     order.UpdatedAt.Unix(),
		UserId:        order.UserId,
		ProductId:     order.ProductId,
		BossId:        order.BossId,
		Num:           order.Num,
		AddressName:   address.Name,
		AddressPhone:  address.Phone,
		Address:       address.Address,
		Type:          order.Type,
		ProductName:   product.Name,
		ImgPath:       conf.R2Domain + "/" + product.ImgPath,
		DiscountPrice: order.Money,
	}
}

func BuildOrders(ctx context.Context, items []*model.Order) (orders []*Order) {
	addressDao := dao.NewAddressDao(ctx)
	productDao := dao.NewProductDao(ctx)
	for _, item := range items {
		address, err := addressDao.FindById(item.AddressId, item.UserId)
		if err != nil {
			continue
		}
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		orders = append(orders, BuildOrder(item, address, product))
	}
	return
}
