package main

import (
	"net/http"
	"solid-todo/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	r.GET("/db-check", func(c *gin.Context) {
		sqlDB, err := db.DB.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection failed"})
			return
		}

		if err = sqlDB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB ping failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "DB is connected"})
	})

	r.Run(":8080") // ポート 8080 で起動
}
