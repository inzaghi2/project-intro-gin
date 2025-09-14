package main

import "github.com/gin-gonic/gin"

func main() {
	//inisialisasi Gin router
	router := gin.Default()
	//MIddleware :Logger
	router.Use(gin.Logger())
	//MIddleware :Recovery
	router.Use(gin.Recovery())
	//Route Definition
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Hello, " + name + "!",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// Di sini Anda dapat melakukan validasi login, misalnya memeriksa di database, dll.
		// Contoh sdehana: Memersa apakah email da Password cocok.
		if loginData.Email == "example@example.com" && loginData.Password == "password123" {
			c.JSON(200, gin.H{
				"message": "Login successful",
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Invalid credentials",
			})
		}
		//Menambhakan endpoint untuk mengambil parameter wuery
		router.GET("/users", func(c *gin.Context) {
			name := c.Query("name")
			if name == "" {
				c.JSON(400, gin.H{
					"error": "Name parameter is missing",
				})
				return
			}
			c.JSON(200, gin.H{
				"message": "hello, " + name + "!",
			})
		})

	})
	//jalankan server
	router.Run(":8080")
}
