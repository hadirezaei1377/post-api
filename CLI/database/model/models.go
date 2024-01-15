package database

import "time"

// Trade represents a trade record
type Trade struct {
	ID      int       `json:"id"`
	ImageID int       `json:"image_id"`
	DateEn  time.Time `json:"date_en"`
	Open    float64   `json:"open"`
	High    float64   `json:"high"`
	Low     float64   `json:"low"`
	Close   float64   `json:"close"`
}

// Image represents an instrument record
type Image struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
