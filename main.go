package main

import "github.com/idantp/momentum-stock-app/internal/db"

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// polygonApiKey := os.Getenv("POLYGON_API_KEY")
	// c := polygon.New(polygonApiKey)
	// log.Printf("Polygon API Key: %s", polygonApiKey)

	// params := models.GetTickerSnapshotParams{
	// 	Ticker:     "AAPL",
	// 	Locale:     "us",
	// 	MarketType: "stocks",
	// }

	// res, err := c.GetTickerSnapshot(context.Background(), &params)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("%+v", res)
	// log.Printf("Ticker: %+v", res.Snapshot.Ticker)
	// log.Printf("Ticker: %+v", res.Snapshot.LastTrade.Price)

	// TODO: consider using pgxpool.Connect(...)
	conn := db.ConnectDB()
	defer db.CloseDB(conn)
	db.CreateStocksTable(conn)
}
