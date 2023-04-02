package filter

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

func Options(c *gin.Context) {
    if strings.Contains(c.Request.Method, "OPTIONS") {
        c.AbortWithStatus(http.StatusOK)
    }
}
