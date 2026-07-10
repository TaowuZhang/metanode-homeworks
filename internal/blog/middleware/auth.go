package middleware

import (
	"net/http"
	"strings"

	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/utils"

	"github.com/gin-gonic/gin"
)

const UserIDKey = "userID"

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			utils.Error(c, http.StatusUnauthorized, 40101, "missing bearer token")
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(strings.TrimPrefix(header, "Bearer "), secret)
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, 40102, "invalid token")
			c.Abort()
			return
		}
		c.Set(UserIDKey, claims.UserID)
		c.Next()
	}
}

func CurrentUserID(c *gin.Context) uint {
	v, ok := c.Get(UserIDKey)
	if !ok {
		return 0
	}
	id, _ := v.(uint)
	return id
}
