package retailoutlet

import (
	httpprovider "my-market-server/main/http_provider"
	retailoutlet_dto "my-market-server/main/retail_outlet/dto"
	"net/http"

	"gorm.io/gorm"
)

func NewRetailOutletController(db *gorm.DB) []httpprovider.IController {
	service := *NewReatilOutletService(&RetailOutletRepo{Repo: db})
	return []httpprovider.IController{
		httpprovider.GenericControllerAdapter[any]{
			Controller: findAll(service),
		},
		httpprovider.GenericControllerAdapter[retailoutlet_dto.CreateRetailOutletDto]{
			Controller: create(service),
		},
	}
}

func findAll(service RetailOutletService) httpprovider.Controller[any] {
	return httpprovider.Controller[any]{
		Path:   "",
		Method: http.MethodGet,
		Handler: func(r httpprovider.HttpResponse[any]) {
			var outlets []RetailOutletModel
			err := service.service.FindAll(&outlets)
			if err != nil {
				r.Exception(http.StatusInternalServerError, "internal server error", nil)
				return
			}
			r.SendJson(outlets)
		},
	}
}

func create(service RetailOutletService) httpprovider.Controller[retailoutlet_dto.CreateRetailOutletDto] {
	return httpprovider.Controller[retailoutlet_dto.CreateRetailOutletDto]{
		Path:   "",
		Method: http.MethodPost,
		Handler: func(r httpprovider.HttpResponse[retailoutlet_dto.CreateRetailOutletDto]) {
			model := service.Create(r.Body)
			r.SendJson(model)
		},
	}
}
