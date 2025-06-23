package retailoutlet

import (
	"context"
	"encoding/json"
	"fmt"
	httpprovider "my-market-server/main/http_provider"
	"net/http"

	"gorm.io/gorm"
)

func Controller(db *gorm.DB) []httpprovider.Controller {
	service := *NewReatilOutletService(&RetailOutletRepo{Repo: db})
	return []httpprovider.Controller{
		findAll(service),
	}
}

func findAll(service RetailOutletService) httpprovider.Controller {
	return httpprovider.Controller{
		Path:   "",
		Method: http.MethodGet,
		Handler: func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			var outlets []RetailOutletModel
			err := service.service.FindAll(&outlets)
			if err != nil {
				fmt.Fprintln(w, "Inter server error")
				return
			}

			w.Header().Set("Content-type", "application/json")
			data, parseErr := json.Marshal(outlets)
			if parseErr != nil {
				fmt.Fprintln(w, "Inter server error")
				return
			}

			w.Write(data)
		},
	}
}
