package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"quote-generator/database"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func main() {
	db := database.GetMongoInstance()
	defer db.Client.Connect(context.Background())
	log.Println("MONGODB CONNECTED")

	router := gin.Default()
	api := router.Group("/api")

	api.POST("/", CreateQuotes)
	api.GET("/", GetQuotes)

	router.Run()
}

func CreateQuotes(ctx *gin.Context) {
	// Get input quotes
	var quotes []Quote
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to read quotes info",
			"error":   true,
		})
		return
	}

	err = json.Unmarshal(body, &quotes)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to read quotes info",
			"error":   true,
		})
		return
	}

	// Create on database
	insert := make([]interface{}, 0)
	for _, q := range quotes {
		insert = append(insert, q)
	}
	collection := database.GetMongoInstance().Db.Collection("quotes")
	_, err = collection.InsertMany(context.Background(), insert)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Fail to create quotes",
			"error":   true,
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "Create quotes succesfully",
		"error":   false,
		"data": gin.H{
			"quotes": quotes,
		},
	})
}

func GetQuotes(ctx *gin.Context) {
	param := ctx.Query("number")
	number, err := strconv.Atoi(param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Wrong parameter",
			"error":   true,
		})
		return
	}


	collection := database.GetMongoInstance().Db.Collection("quotes")

	var quotes []Quote
	pipeline := []interface{}{
		bson.M{
			"$sample": bson.M{
				"size": number,
			},
		},
	}
	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to get quotes",
			"error":   true,
		})
		return
	}

	if err = cursor.All(context.Background(), &quotes); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Fail to get quotes",
			"error":   true,
		})
		return
	}

	log.Println(quotes)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get random quotes successfully",
		"error":   false,
		"data": gin.H{
			"quotes": quotes,
		},
	})
}
