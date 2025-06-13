package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// LoadRecords reads GO_test_5m.csv and returns []Record
func LoadRecords(path string) ([]Record, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	header, err := r.Read() // header row
	if err != nil {
		return nil, err
	}

	// Column index map
	idx := map[string]int{}
	for i, col := range header {
		idx[strings.TrimSpace(col)] = i
	}
	// required columns
	req := []string{
		"transaction_id", "transaction_date", "user_id", "country", "region",
		"product_id", "product_name", "category", "price", "quantity",
		"total_price", "stock_quantity", "added_date",
	}
	for _, col := range req {
		if _, ok := idx[col]; !ok {
			return nil, fmt.Errorf("column %q not found", col)
		}
	}

	var recs []Record
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		// Parse dates (format "2006-01-02" or adjust if includes time)
		td, err := time.Parse("2006-01-02", row[idx["transaction_date"]])
		if err != nil {
			// try datetime format if needed:
			td, err = time.Parse("2006-01-02 15:04:05", row[idx["transaction_date"]])
			if err != nil {
				// skip or set zero
				td = time.Time{}
			}
		}
		ad, err := time.Parse("2006-01-02", row[idx["added_date"]])
		if err != nil {
			ad = time.Time{}
		}

		price, _ := strconv.ParseFloat(row[idx["price"]], 64)
		qty, _ := strconv.Atoi(row[idx["quantity"]])
		tp, _ := strconv.ParseFloat(row[idx["total_price"]], 64)
		sq, _ := strconv.Atoi(row[idx["stock_quantity"]])

		recs = append(recs, Record{
			TransactionID:   row[idx["transaction_id"]],
			TransactionDate: td,
			UserID:          row[idx["user_id"]],
			Country:         row[idx["country"]],
			Region:          row[idx["region"]],
			ProductID:       row[idx["product_id"]],
			ProductName:     row[idx["product_name"]],
			Category:        row[idx["category"]],
			Price:           price,
			Quantity:        qty,
			TotalPrice:      tp,
			StockQuantity:   sq,
			AddedDate:       ad,
		})
	}

	return recs, nil
}
