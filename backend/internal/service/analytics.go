package service

import (
	"sort"
	"time"

	"github.com/Chethana-it/qoria-go-dashboard/backend/internal/data"
)

// RevenueByCountry sums total_price per country.
func RevenueByCountry(records []data.Record) map[string]float64 {
	totals := make(map[string]float64, len(records))
	for _, r := range records {
		totals[r.Country] += r.TotalPrice
	}
	return totals
}

// ProductStats holds aggregated stats for a product.
type ProductStats struct {
	ProductID     string
	ProductName   string
	QuantitySold  int
	StockQuantity int
}

// CountryProductStats holds stats per country×product.
type CountryProductStats struct {
	Country          string  `json:"country"`
	ProductName      string  `json:"product_name"`
	TotalRevenue     float64 `json:"total_revenue"`
	TransactionCount int     `json:"transaction_count"`
}

// CountryProductTable computes total revenue and transaction count
// for each country×product combination, and returns a slice sorted
// descending by TotalRevenue.
func CountryProductTable(records []data.Record) []CountryProductStats {
	// key = country + "|" + productName
	agg := make(map[string]*CountryProductStats, len(records))
	for _, r := range records {
		key := r.Country + "|" + r.ProductName
		if ps, ok := agg[key]; ok {
			ps.TotalRevenue += r.TotalPrice
			ps.TransactionCount++
		} else {
			agg[key] = &CountryProductStats{
				Country:          r.Country,
				ProductName:      r.ProductName,
				TotalRevenue:     r.TotalPrice,
				TransactionCount: 1,
			}
		}
	}
	// to slice
	out := make([]CountryProductStats, 0, len(agg))
	for _, v := range agg {
		out = append(out, *v)
	}
	// sort desc by revenue
	sort.Slice(out, func(i, j int) bool {
		return out[i].TotalRevenue > out[j].TotalRevenue
	})
	return out
}

// TopProducts returns the top N products by quantity sold,
// including the latest stock quantity seen.
func TopProducts(records []data.Record, topN int) []ProductStats {
	agg := make(map[string]*ProductStats, len(records))
	for _, r := range records {
		if ps, ok := agg[r.ProductID]; ok {
			ps.QuantitySold += r.Quantity
			ps.StockQuantity = r.StockQuantity // overwrite; last seen
		} else {
			agg[r.ProductID] = &ProductStats{
				ProductID:     r.ProductID,
				ProductName:   r.ProductName,
				QuantitySold:  r.Quantity,
				StockQuantity: r.StockQuantity,
			}
		}
	}

	// move to slice and sort descending by QuantitySold
	slice := make([]ProductStats, 0, len(agg))
	for _, ps := range agg {
		slice = append(slice, *ps)
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].QuantitySold > slice[j].QuantitySold
	})

	if len(slice) > topN {
		return slice[:topN]
	}
	return slice
}

// MonthStats holds total quantity sold in a specific month.
type MonthStats struct {
	Month        string // e.g. "2023-07"
	QuantitySold int
}

// SalesByMonth aggregates total quantity sold per calendar month.
func SalesByMonth(records []data.Record) []MonthStats {
	agg := make(map[string]int, len(records))
	for _, r := range records {
		month := r.TransactionDate.Format("2006-01")
		agg[month] += r.Quantity
	}

	// convert to slice and sort by month ascending
	slice := make([]MonthStats, 0, len(agg))
	for m, qty := range agg {
		slice = append(slice, MonthStats{Month: m, QuantitySold: qty})
	}
	sort.Slice(slice, func(i, j int) bool {
		ti, _ := time.Parse("2006-01", slice[i].Month)
		tj, _ := time.Parse("2006-01", slice[j].Month)
		return ti.Before(tj)
	})
	return slice
}

// RegionStats holds revenue and quantity for a region.
type RegionStats struct {
	Region       string
	TotalRevenue float64
	QuantitySold int
}

// RevenueByRegion returns the top N regions by total revenue,
// including the total items sold.
func RevenueByRegion(records []data.Record, topN int) []RegionStats {
	agg := make(map[string]*RegionStats, len(records))
	for _, r := range records {
		if rs, ok := agg[r.Region]; ok {
			rs.TotalRevenue += r.TotalPrice
			rs.QuantitySold += r.Quantity
		} else {
			agg[r.Region] = &RegionStats{
				Region:       r.Region,
				TotalRevenue: r.TotalPrice,
				QuantitySold: r.Quantity,
			}
		}
	}

	// to slice and sort by TotalRevenue desc
	slice := make([]RegionStats, 0, len(agg))
	for _, rs := range agg {
		slice = append(slice, *rs)
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i].TotalRevenue > slice[j].TotalRevenue
	})

	if len(slice) > topN {
		return slice[:topN]
	}
	return slice
}
