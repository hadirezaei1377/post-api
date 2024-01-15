package CLI

import (
	"database/sql"
	"flag"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Parse command line arguments
	numRecords := flag.Int("num-records", 1000, "Number of records to generate")
	flag.Parse()

	// Connect to the database
	db, err := sql.Open("postgres", "user=username dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate and insert random trade records
	for i := 0; i < *numRecords; i++ {
		Trade := trade,
		trade := Trade{
			ImageID: rand.Intn(2) + 1,
			DateEn:  time.Now(),
			Open:    rand.Float64() * 1000,
			High:    rand.Float64() * 1000,
			Low:     rand.Float64() * 1000,
			Close:   rand.Float64() * 1000,
		}

		_, err := db.Exec("INSERT INTO Trade (ImageID, DateEn, Open, High, Low, Close) VALUES ($1, $2, $3, $4, $5, $6)",
			trade.ImageID, trade.DateEn, trade.Open, trade.High, trade.Low, trade.Close)

		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("%d records inserted successfully.\n", *numRecords)
}
