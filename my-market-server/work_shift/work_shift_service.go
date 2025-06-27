package workshift

import (
	"fmt"
	retailoutlet "my-market-server/main/retail_outlet"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type WorkShiftService struct {
	Repo                WorkShiftRepo
	RetailOutletService retailoutlet.RetailOutletService
}

func NewWorkShiftService(repo WorkShiftRepo, retailOutletService retailoutlet.RetailOutletService) WorkShiftService {
	return WorkShiftService{Repo: repo, RetailOutletService: retailOutletService}
}

func (s *WorkShiftService) FindAll(query url.Values) ([]WorkShiftModel, error) {
	limit := 10
	order := "asc"

	if limitArr, ok := query["limit"]; ok {
		value, err := strconv.Atoi(limitArr[0])
		if err == nil {
			limit = value
		}
	}

	if orderArr, ok := query["order"]; ok {
		od := strings.ToLower(orderArr[0])
		if od == "asc" || od == "desc" {
			order = od
		}
	}

	return s.Repo.FindAll(limit, order)
}

func (s *WorkShiftService) FindById(id int) (WorkShiftModel, error) {
	return s.Repo.FindById(id)
}

func (s *WorkShiftService) Create(dto CreateWorkShiftDto) (WorkShiftModel, error) {
	_, rtErr := s.RetailOutletService.FindById(dto.RetailOutletId)
	if rtErr != nil {
		return WorkShiftModel{}, fmt.Errorf("retail outlet [%d] not found", dto.RetailOutletId)
	}
	model := WorkShiftModel{
		CreatedAt:      time.Now(),
		Cash:           dto.Cash,
		Cashless:       dto.Cashless,
		CashRegister:   dto.CashRegister,
		RetailOutletId: dto.RetailOutletId,
	}

	if dto.CardTransfers != nil {
		for _, transfer := range *dto.CardTransfers {
			transferModel := CardTransferModel{Sum: transfer}
			model.CardTransfers = append(model.CardTransfers, transferModel)
		}
	}

	if dto.Expenses != nil {
		for _, expense := range *dto.Expenses {
			expenseModel := ExpenseModel{
				Article: expense.Article,
				Debit:   expense.Debit,
				Credit:  expense.Credit,
				Payed:   expense.Payed,
			}
			model.Expenses = append(model.Expenses, expenseModel)
		}
	}

	return s.Repo.Create(model)
}
