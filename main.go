package main

import (
	"net/http"
	"context"
	"io/ioutil"
	"encoding/json"
	"math/rand"
	"log"
	"time"

	"github.com/gin-gonic/gin"
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

	api.POST("/", CreateManyQuotes)
	api.GET("/", GetQuote)

	router.Run()
}



func CreateManyQuotes(ctx *gin.Context) {
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

func GetQuote(ctx *gin.Context) {
	collection := database.GetMongoInstance().Db.Collection("quotes")
	cursor, _ := collection.Find(context.Background(), gin.H{})
	var quotes []Quote 
	if err := cursor.All(context.Background(), &quotes); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Fail to read quotes info",
			"error":   true,
		})
		return
	}

	// Random index of quote in slice
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(quotes) - 1)
	quote := quotes[index]
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get random quote successfully",
		"error": false,
		"data": gin.H{
			"quote": quote,
		},
	})
}
