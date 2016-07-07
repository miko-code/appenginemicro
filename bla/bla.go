package bla

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love bla")
}
func init() {
	r := gin.New()

	r.Handle("GET", "/bla", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "bla",
		})
	})

	http.Handle("/bla", r)
	//	http.HandleFunc("/bla", handler)

}
