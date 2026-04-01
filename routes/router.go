package routes

import (
	api "gin_mall/api/v1"
	"gin_mall/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("static"))
	v1 := r.Group("/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "pong")
		})
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		v1.GET("carousels", api.ListCarousel)

		v1.GET("products", api.ListProduct)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			authed.POST("money", api.ShowMoney)

			authed.POST("product", api.CreateProduct)
			authed.POST("products", api.SearchProduct)
		}

	}
	return r
}
