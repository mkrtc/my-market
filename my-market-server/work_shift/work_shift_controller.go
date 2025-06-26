package workshift

import (
	hc "my-market-server/main/http_client"
	retailoutlet "my-market-server/main/retail_outlet"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func NewWorkShiftController(db *gorm.DB) []hc.IController {
	service := *NewWorkShiftService(WorkShiftRepo{Repo: db}, *retailoutlet.NewRetailOutletService(&retailoutlet.RetailOutletRepo{Repo: db}))
	return []hc.IController{
		hc.GenericControllerAdapter[any]{Controller: findById(service)},
		hc.GenericControllerAdapter[any]{Controller: findAll(service)},
		hc.GenericControllerAdapter[CreateWorkShiftDto]{Controller: create(service)},
	}
}

func findAll(service WorkShiftService) hc.Controller[any] {
	return hc.Controller[any]{
		Path:   "",
		Method: http.MethodGet,
		Handler: func(r hc.HttpResponse[any]) {

			outlets, err := service.FindAll(r.Query)
			if err != nil {
				r.Exception(http.StatusInternalServerError, "internal server error", nil)
				return
			}
			r.SendJson(outlets)
		},
	}
}

func findById(service WorkShiftService) hc.Controller[any] {
	return hc.Controller[any]{
		Path:   "/by-id",
		Method: http.MethodGet,
		Handler: func(r hc.HttpResponse[any]) {
			id := 0
			if idArr, ok := r.Query["id"]; ok {
				value, err := strconv.Atoi(idArr[0])
				if err == nil {
					id = value
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

func create(service WorkShiftService) hc.Controller[CreateWorkShiftDto] {
	return hc.Controller[CreateWorkShiftDto]{
		Path:   "",
		Method: http.MethodPost,
		Handler: func(r hc.HttpResponse[CreateWorkShiftDto]) {
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
