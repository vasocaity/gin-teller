package main

import (
	"go-2025/cache"
	"go-2025/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// 🔹 โหลดค่าจากไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	cache.Init()
	db.InitDB()

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	r.Run()
}
