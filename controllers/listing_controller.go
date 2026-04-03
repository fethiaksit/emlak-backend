package controllers

import (
	"context"
	"net/http"
	"time"

	"emlak-backend/config"
	"emlak-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collectionName = "listings"

func isValidRisk(risk string) bool {
	return risk == "düşük" || risk == "orta" || risk == "yüksek"
}

func GetListings(c *gin.Context) {
	collection := config.DB.Collection(collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "İlanlar alınamadı"})
		return
	}
	defer cursor.Close(ctx)

	var listings []models.Listing
	if err := cursor.All(ctx, &listings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "İlanlar okunamadı"})
		return
	}

	c.JSON(http.StatusOK, listings)
}

func GetListingByID(c *gin.Context) {
	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ilan id"})
		return
	}

	collection := config.DB.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var listing models.Listing
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&listing)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "İlan bulunamadı"})
		return
	}

	c.JSON(http.StatusOK, listing)
}

func CreateListing(c *gin.Context) {
	var listing models.Listing
	if err := c.ShouldBindJSON(&listing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz JSON body"})
		return
	}

	if listing.Title == "" || listing.Description == "" || listing.Price <= 0 || listing.City == "" || listing.District == "" || listing.PropertyType == "" || listing.ListingType == "" || listing.EarthquakeRisk == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Zorunlu alanlar eksik veya geçersiz"})
		return
	}

	if !isValidRisk(listing.EarthquakeRisk) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "earthquakeRisk yalnızca düşük, orta veya yüksek olabilir"})
		return
	}

	listing.ID = primitive.NewObjectID()
	listing.CreatedAt = time.Now()

	collection := config.DB.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, listing)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "İlan eklenemedi"})
		return
	}

	c.JSON(http.StatusCreated, listing)
}

func DeleteListing(c *gin.Context) {
	idParam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ilan id"})
		return
	}

	collection := config.DB.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "İlan silinemedi"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "İlan bulunamadı"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "İlan silindi"})
}
