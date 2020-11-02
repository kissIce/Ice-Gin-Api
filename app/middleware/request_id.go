package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	//"ice/global"
	//"strconv"
)

func RequestId() gin.HandlerFunc  {
	return func(c *gin.Context) {
		if c.Request.Header.Get("X-Request-Id") == "" {
			c.Writer.Header().Set("X-Request-Id", uuid.Must(uuid.NewV4()).String())
			//c.Writer.Header().Set("X-Request-Id", strconv.FormatInt(global.IceSnowflake.GetId(), 10))
		}
		c.Next()
	}
}
