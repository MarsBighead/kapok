package api

import (
	"github.com/gin-gonic/gin"
)

//Service service layer for API
type Service interface {
	Add(c *gin.Context)
	Update(c *gin.Context)
	Query(c *gin.Context)
	Del(c *gin.Context)
}
