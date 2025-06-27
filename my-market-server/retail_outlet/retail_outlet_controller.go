package retailoutlet

import (
	hc "my-market-server/main/http_client"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func NewRetailOutletController(db *gorm.DB) []hc.IController {
	service := NewRetailOutletService(&RetailOutletRepo{Repo: db})
	return []hc.IController{
		hc.GenericControllerAdapter[any]{
			Controller: findAll(service),
		},
		hc.GenericControllerAdapter[any]{
			Controller: findById(service),
		},
		hc.GenericControllerAdapter[CreateRetailOutletDto]{
			Controller: create(service),
		},
	}
}

func findAll(service RetailOutletService) hc.Controller[any] {
	return hc.Controller[any]{
		Path:   "",
		Method: http.MethodGet,
		Handler: func(r hc.HttpResponse[any]) {
			outlets, err := service.FindAll()
			if err != nil {
				r.Exception(http.StatusInternalServerError, "internal server error", nil)
				return
			}
			r.SendJson(outlets)
		},
	}
}

func findById(service RetailOutletService) hc.Controller[any] {
	return hc.Controller[any]{
		Path:   "/by-id",
		Method: http.MethodGet,
		Handler: func(r hc.HttpResponse[any]) {
			id := 0
			if i, ok := r.Query["id"]; ok {
				if len(i) > 0 {
					value, err := strconv.Atoi(i[0])
					if err == nil {
						id = value
					}
				}
			}

			if id == 0 {
				r.Exception(http.StatusBadRequest, "Id is required", nil)
				return
			}
			model, err := service.FindById(id)
			if err != nil {
				r.Exception(http.StatusBadRequest, err.Error(), nil)
				return
			}

			r.SendJson(model)

		},
	}
}

func create(service RetailOutletService) hc.Controller[CreateRetailOutletDto] {
	return hc.Controller[CreateRetailOutletDto]{
		Path:   "",
		Method: http.MethodPost,
		Handler: func(r hc.HttpResponse[CreateRetailOutletDto]) {
			if !r.HasBody {
				r.Exception(http.StatusBadRequest, "Incorrect body!", nil)
				return
			}
			model, err := service.Create(r.Body)
			if err != nil {
				r.Exception(http.StatusBadRequest, err.Error(), nil)
				return
			}

			r.SendJson(model)
		},
	}
}
