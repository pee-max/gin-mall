package middleware

import (
	"fmt"
	"gin_mall/pkg/e"
	"gin_mall/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.Success
		token := c.GetHeader("Authorization")
		if token == "" {
			code = e.ErrorAuthTokenEmpty
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
				fmt.Println(err, token)
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthTokenTimeOut
			}
		}
		if code != e.Success {
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
