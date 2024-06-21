package routes

import (
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	Products(router)
}
