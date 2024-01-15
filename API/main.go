package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// represent a trade record
type Trade struct {
	ID           int       `json:"id"`
	ImageID int       `json:"image_id"`
	DateEn       time.Time `json:"date_en"`
	Open         float64   `json:"open"`
	High         float64   `json:"high"`
	Low          float64   `json:"low"`
	Close        float64   `json:"close"`
}

// represent an instrument record
type Image struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	db, err := sql.Open("postgres", "user=username dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default() // Initialize the Gin router 

	// a route for the API endpoint
	router.GET("/api/last_trade", func(c *gin.Context) {
		lastTrades := []Trade{
			{ID: 2, ImageID: 1, DateEn: time.Now(), Open: 1002, High: 2002, Low: 302, Close: 402},
			{ID: 5, ImageID: 2, DateEn: time.Now(), Open: 1005, High: 2005, Low: 305, Close: 405},
		}

		c.JSON(http.StatusOK, lastTrades)
	})


	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}