package test

/*
	@author: Allen
	@since: 2023/3/22
	@desc: //TODO
*/
import (
	"github.com/Allen9012/gee/gee"
	"net/http"
	"testing"
)

func TestGee(t *testing.T) {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
