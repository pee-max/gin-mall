package serializer

import (
	"gin_mall/model"
	"gin_mall/pkg/util"
)

type Money struct {
	UserID    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(item *model.User, key string) (Money, error) {
	util.Encrypt.SetKey(key)
	moneyStr, err := util.Encrypt.AesDecoding(item.Money)
	if err != nil {
		return Money{}, err
	}
	return Money{
		UserID:    item.ID,
		UserName:  item.UserName,
		UserMoney: moneyStr,
	}, nil
}
