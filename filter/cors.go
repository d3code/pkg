package filter

import (
    "github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
    c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Add("Access-Control-Allow-Methods", "POST, PUT, GET, OPTIONS, DELETE")
    c.Writer.Header().Add("Access-Control-Allow-Headers", "Authorization, Content-Type")
    c.Writer.Header().Add("Access-Control-Max-Age", "3600")

    c.Next()
}
