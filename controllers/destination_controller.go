package controllers

import (
	"TouristAppSolution/api/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var destinationCollection *mongo.Collection

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Greška prilikom učitavanja .env datoteke: %v", err)
	}

	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	clientOptions := options.Client().ApplyURI(connectionString)

	client, error := mongo.Connect(context.TODO(), clientOptions)

	if error != nil {
		log.Fatalf("Greška pri povezivanju sa MongoDB-om: %v", error)
	}

	fmt.Println("MongoDB connection success")

	dbName := os.Getenv("DBNAME")
    	colName := os.Getenv("COLNAME")

	destinationCollection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

func GetDestinationsByCity(c *gin.Context) {
    city := c.Param("city")

    filter := bson.M{"City": city}
    cursor, err := destinationCollection.Find(context.Background(), filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var destinations []models.Destination
    if err := cursor.All(context.Background(), &destinations); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, destinations)
}

func GetDestinationsByLandmarks(c *gin.Context) {
    landmarks := c.Query("landmark")

    filter := bson.M{"Landmarks": landmarks}
    cursor, err := destinationCollection.Find(context.Background(), filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var destinations []models.Destination
    if err := cursor.All(context.Background(), &destinations); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, destinations)
}