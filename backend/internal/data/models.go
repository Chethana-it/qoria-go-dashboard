package data

import "time"

// Record represents one row from GO_test_5m.csv
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
