package middleware

import (
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	CORSMiddleware() gin.HandlerFunc
	TracerMiddleware() gin.HandlerFunc
}
