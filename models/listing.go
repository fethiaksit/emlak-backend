package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Listing struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title          string             `bson:"title" json:"title"`
	Description    string             `bson:"description" json:"description"`
	Price          float64            `bson:"price" json:"price"`
	City           string             `bson:"city" json:"city"`
	District       string             `bson:"district" json:"district"`
	PropertyType   string             `bson:"propertyType" json:"propertyType"`
	ListingType    string             `bson:"listingType" json:"listingType"`
	EarthquakeRisk string             `bson:"earthquakeRisk" json:"earthquakeRisk"`
	CreatedAt      time.Time          `bson:"createdAt" json:"createdAt"`
}
