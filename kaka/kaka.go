package kaka

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love kaka")
}
func init() {
	r := gin.New()

	r.Handle("GET", "/kaka", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "kaka",
		})
	})

	http.Handle("/kaka", r)

	//	http.HandleFunc("/kaka", handler)

}
