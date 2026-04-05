package v1

import (
	"gin_mall/pkg/util"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListFavorites(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	listFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&listFavoriteService); err == nil {
		res := listFavoriteService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Infoln(err)
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func CreateFavorites(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&createFavoriteService); err == nil {
		res := createFavoriteService.Creat(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Infoln(err)
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}

func DeleteFavorites(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	deleteFavoriteService := service.FavoriteService{}
	if err := c.ShouldBind(&deleteFavoriteService); err == nil {
		res := deleteFavoriteService.Delete(c.Request.Context(), claim.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		util.LogrusObj.Infoln(err)
		c.JSON(http.StatusOK, ErrorResponse(err))
	}
}
