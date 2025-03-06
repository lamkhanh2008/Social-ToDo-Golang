package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Tạo một router mới
	r := gin.Default()

	// Định nghĩa một route cho GET request
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Định nghĩa một route cho POST request
	r.POST("/create", func(c *gin.Context) {
		var json map[string]interface{}

		// Bind JSON request body vào biến json
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}

		// Trả về response với dữ liệu đã nhận
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data":   json,
		})
	})

	// Khởi động server ở cổng 8080
	r.Run(":3000")
}
