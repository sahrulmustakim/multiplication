package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	page := gin.Default()
	theme := template.Must(template.New("main").ParseGlob("templates/*"))

	page.SetHTMLTemplate(theme)
	page.GET("/", func(result *gin.Context) {
		data := []string{}
		var x, y = 1, 1
		for x <= 9 {
			for y <= 9 {
				data = append(data, fmt.Sprintf("%v x %v = %v", x, y, x*y))
				y++
			}
			y = 1
			x++
		}

		result.HTML(http.StatusOK, "content.tmpl", gin.H{"data": data})
	})

	log.Fatal(page.Run(":" + port))
}
