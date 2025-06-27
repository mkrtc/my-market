package retailoutlet

import (
	hc "my-market-server/main/http_client"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func NewSeoController(db *gorm.DB) []hc.IController {
	service := NewSeoService(&SeoRepo{Repo: db})
	return []hc.IController{
		hc.GenericControllerAdapter[any]{Controller: findSeoById(service)},
		hc.GenericControllerAdapter[any]{Controller: findAllSeo(service)},
		hc.GenericControllerAdapter[CreateSeoDto]{Controller: createSeo(service)},
	}
}

func findSeoById(service SeoService) hc.Controller[any] {
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

func findAllSeo(service SeoService) hc.Controller[any] {
	return hc.Controller[any]{
		Path:   "",
		Method: http.MethodGet,
		Handler: func(r hc.HttpResponse[any]) {
			model, err := service.FindAll()
			if err != nil {
				r.Exception(http.StatusBadRequest, err.Error(), nil)
				return
			}
			r.SendJson(model)
		},
	}
}

func createSeo(service SeoService) hc.Controller[CreateSeoDto] {
	return hc.Controller[CreateSeoDto]{
		Path:   "",
		Method: http.MethodPost,
		Handler: func(r hc.HttpResponse[CreateSeoDto]) {
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
