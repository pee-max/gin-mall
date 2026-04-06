package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
	"strconv"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `form:"address" json:"address"`
}

func (service *AddressService) Create(ctx context.Context, uid uint) serializer.Response {
	var address *model.Address
	addressDao := dao.NewAddressDao(ctx)
	code := e.Success
	address = &model.Address{
		UserID:  uid,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.Create(address)
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

func (service *AddressService) Get(ctx context.Context, id string, uid uint) serializer.Response {
	var address *model.Address
	addressDao := dao.NewAddressDao(ctx)
	code := e.Success
	aid, _ := strconv.Atoi(id)
	address, err := addressDao.FindById(uint(aid), uid)
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
		Data:   serializer.BuildAddress(address),
	}
}

func (service *AddressService) List(ctx context.Context, uid uint) serializer.Response {
	var addresses []*model.Address
	addressDao := dao.NewAddressDao(ctx)
	code := e.Success
	addresses, err := addressDao.List(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildAddresses(addresses), uint(len(addresses)))
}

func (service *AddressService) Update(ctx context.Context, uid uint, id string) serializer.Response {
	var address *model.Address
	addressDao := dao.NewAddressDao(ctx)
	code := e.Success
	aid, _ := strconv.Atoi(id)
	address = &model.Address{
		UserID:  uid,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.Update(address, uint(aid), uid)
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

func (service *AddressService) Delete(ctx context.Context, uid uint, id string) serializer.Response {
	addressDao := dao.NewAddressDao(ctx)
	code := e.Success
	aid, _ := strconv.Atoi(id)
	err := addressDao.Delete(uint(aid), uint(uid))
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
