package service

import (
	"time"
	"uuid"

	"github.com/rajat-mangla/go-boilerplate/domain"
)

type CreateRepaymentPlanInput struct {
	CustomerID        string
	CommodityCost     float64
	CommodityCategory string
	Installments      int
	PromoCode         string
	PurchaseDate      *time.Time
}

type RepaymentPlanResult struct {
	ID   string
	Plan *domain.RepaymentPlan
}

type RepaymentPlanService struct{}

func NewRepaymentPlanService() *RepaymentPlanService {
	return &RepaymentPlanService{}
}

func (s *RepaymentPlanService) CreateRepaymentPlan(input CreateRepaymentPlanInput) (*RepaymentPlanResult, error) {
	purchaseDate := time.Now()
	if input.PurchaseDate != nil {
		purchaseDate = *input.PurchaseDate
	}

	plan, err := domain.Compute(domain.ComputeInput{
		CustomerID:        input.CustomerID,
		CommodityCost:     input.CommodityCost,
		CommodityCategory: input.CommodityCategory,
		Installments:      input.Installments,
		PromoCode:         input.PromoCode,
		PurchaseDate:      purchaseDate,
	})
	if err != nil {
		return nil, err
	}

	return &RepaymentPlanResult{
		ID:   uuid.NewV7().String(),
		Plan: plan,
	}, nil
}
