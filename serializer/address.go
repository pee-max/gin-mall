package serializer

import "gin_mall/model"

type Address struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Seen     bool   `json:"seen"`
	CreateAt int64  `json:"create_at"`
}

func BuildAddress(item *model.Address) Address {
	return Address{
		Id:       item.ID,
		UserId:   item.UserID,
		Name:     item.Name,
		Phone:    item.Phone,
		Address:  item.Address,
		CreateAt: item.CreatedAt.Unix(),
	}
}

func BuildAddresses(items []*model.Address) (address []Address) {
	for _, item := range items {
		address = append(address, BuildAddress(item))
	}
	return
}
