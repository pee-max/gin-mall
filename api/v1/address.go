package v1

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAddress(c *gin.Context) {
	createAddressService := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createAddressService); err == nil {
		res := createAddressService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
func GetAddress(c *gin.Context) {
	getAddressService := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&getAddressService); err == nil {
		res := getAddressService.Get(c.Request.Context(), c.Param("id"), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
func ListAddress(c *gin.Context) {
	listAddressService := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listAddressService); err == nil {
		res := listAddressService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
func UpdateAddress(c *gin.Context) {
	updateAddressService := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateAddressService); err == nil {
		res := updateAddressService.Update(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
func DeleteAddress(c *gin.Context) {
	deleteAddressService := service.AddressService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteAddressService); err == nil {
		res := deleteAddressService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
