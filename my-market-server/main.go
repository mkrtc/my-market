package main

import (
	"my-market-server/main/database"
	httpprovider "my-market-server/main/http_provider"
	retailoutlet "my-market-server/main/retail_outlet"
	workshift "my-market-server/main/work_shift"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	db := database.Init()

	db.AutoMigrate(
		&workshift.WorkShiftModel{}, &workshift.CardTransferModel{}, &workshift.ExpenseModel{},
		&retailoutlet.RetailOutletModel{}, &retailoutlet.SeoModel{},
	)

	httpClient := httpprovider.NewClient()

	httpClient.RegisterController("/reatil-outlet", retailoutlet.Controller(db)...)

	port := os.Getenv("PORT")
	if port == "" {
		panic("Port is not definde")
	}

	httpClient.Listen(port)
}
