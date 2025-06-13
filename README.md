# Backend Setup Guide

1. **Clone the repository**

   ```bash
   git clone https://github.com/Chethana-it/qoria-go-dashboard.git
   cd qoria-go-dashboard/backend
   ```

2. **Initialize Go module**

   ```bash
   go mod tidy
   ```

   * Installs all dependencies listed in `go.mod`.

3. **Configure environment (CORS)**

   * If you need to call APIs from a different origin, install and enable CORS middleware:

     ```bash
     go get github.com/gin-contrib/cors
     ```
   * In `backend/cmd/server/main.go`, add below after `router := gin.Default()`:

     ```go
     import "github.com/gin-contrib/cors"
     router.Use(cors.Default())
     ```

4. **Place your CSV file**

   * Copy `GO_test_5m.csv` into `backend/data/`:

     ```bash
     mkdir -p data
     cp /path/to/GO_test_5m.csv data/
     ```

5. **Run the server**

   ```bash
   cd cmd/server
   go run main.go
   ```

   * Server listens on `http://localhost:8080`.

6. **Verify endpoints**

   * Health check:  `curl http://localhost:8080/health`
   * Revenue by country:  `curl http://localhost:8080/api/v1/revenue-by-country`
   * Country×Product stats (paginated):

     ```bash
     curl "http://localhost:8080/api/v1/country-product-stats?limit=100&offset=0"
     ```
   * Top products:  `curl "http://localhost:8080/api/v1/top-products?limit=20"`
   * Monthly sales:  `curl http://localhost:8080/api/v1/sales-by-month`
   * Region revenue:  `curl "http://localhost:8080/api/v1/revenue-by-region?limit=30"`

---

# Frontend Setup Guide

1. **Navigate to the frontend folder**

   ```bash
   cd ../frontend
   ```

2. **Install dependencies**

   ```bash
   npm install
   ```

   * Installs React, axios, recharts, etc.

3. **Configure proxy (optional)**

   * In `frontend/package.json`, add:

     ```json
     "proxy": "http://localhost:8080",
     ```
   * Allows relative API calls without CORS headers.

4. **Fix file watcher limits (Linux)**

   * If you see `ENOSPC` errors, create `.env`:

     ```
     CHOKIDAR_USEPOLLING=true
     CHOKIDAR_IGNORE=node_modules/**
     ```

5. **Start the development server**

   ```bash
   npm start
   ```

   * Opens `http://localhost:3000` in your browser.

6. **Verify dashboard**

   * Check that tables and charts load:

     * Country×Product table
     * Top Products bar chart
     * Monthly Sales line chart
     * Top Regions bar chart

7. **Build for production**

   ```bash
   npm run build
   ```

   



