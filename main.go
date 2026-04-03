package main

import (
	"log"
	"os"

	"emlak-backend/config"
	"emlak-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env dosyası yüklenemedi, sistem değişkenleri kullanılacak")
	}

	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Veritabanı bağlantı hatası: %v", err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	routes.ListingRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Sunucu çalışıyor: :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Sunucu başlatılamadı: %v", err)
	}
}
