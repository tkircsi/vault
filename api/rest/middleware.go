package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidContentType = errors.New("invalid content-type header")
)

func JSONContent() gin.HandlerFunc {
	return func(c *gin.Context) {
		ct := c.GetHeader("Content-Type")
		if ct != "application/json" {
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrInvalidContentType.Error()})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
