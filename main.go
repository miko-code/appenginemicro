package appenginemicro

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.New()

	r.Handle("GET", "/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "default",
		})
	})

	http.Handle("/", r)

}
