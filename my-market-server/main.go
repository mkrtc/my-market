package main

import (
	"my-market-server/main/database"
	httpprovider "my-market-server/main/http_provider"
	retailoutlet "my-market-server/main/retail_outlet"
	retailoutlet_models "my-market-server/main/retail_outlet/models"
	workshift "my-market-server/main/work_shift"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	db := database.Init()

	db.AutoMigrate(
		&workshift.WorkShiftModel{}, &workshift.CardTransferModel{}, &workshift.ExpenseModel{},
		&retailoutlet_models.RetailOutletModel{}, &retailoutlet_models.SeoModel{},
	)

	httpClient := httpprovider.NewClient()

	httpClient.RegisterController("/retail-outlet", retailoutlet.NewRetailOutletController(db)...)

	port := os.Getenv("PORT")
	if port == "" {
		panic("Port is not defined")
	}

	httpClient.Listen(port)
}
