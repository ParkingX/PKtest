package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type student struct {
	Name string
	Age  int8
}

func NewGin() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.html", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.Run() 
}
