package data

import "time"

type Record struct {
	TransactionID   string
	TransactionDate time.Time
	UserID          string
	Country         string
	Region          string
	ProductID       string
	ProductName     string
	Category        string
	Price           float64
	Quantity        int
	TotalPrice      float64
	StockQuantity   int
	AddedDate       time.Time
}
