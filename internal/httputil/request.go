package httputil

import "github.com/gin-gonic/gin"

type ContextValues struct {
	UserID int64
}

func ExtractValuesFromContext(c *gin.Context) *ContextValues {
	return &ContextValues{
		UserID: 10, //c.GetInt64("userID"),
	}
}
