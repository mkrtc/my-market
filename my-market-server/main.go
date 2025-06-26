package main

import (
	"my-market-server/main/database"
	hc "my-market-server/main/http_client"
	rt "my-market-server/main/retail_outlet"
	ws "my-market-server/main/work_shift"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	db := database.Init()

	db.AutoMigrate(
		&ws.WorkShiftModel{}, &ws.CardTransferModel{}, &ws.ExpenseModel{},
		&rt.RetailOutletModel{}, &rt.SeoModel{},
	)

	httpClient := hc.NewClient()

	httpClient.RegisterController("/retail-outlet", rt.NewRetailOutletController(db)...)
	httpClient.RegisterController("/seo", rt.NewSeoController(db)...)
	httpClient.RegisterController("/work-shift", ws.NewWorkShiftController(db)...)

	port := os.Getenv("PORT")
	if port == "" {
		panic("Port is not defined")
	}

	httpClient.Listen(port)
}
