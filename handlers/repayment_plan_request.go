package handlers

import (
	"math"
	"time"

	apperrors "github.com/rajat-mangla/go-boilerplate/errors"
	"github.com/rajat-mangla/go-boilerplate/service"
)

const (
	purchaseDateLayout = "2006-01-02"

	MinInstallments = 2
	MaxInstallments = 12
)

type CreateRepaymentPlanRequest struct {
	CustomerID        string  `json:"customer_id"`
	CommodityCost     float64 `json:"commodity_cost"`
	CommodityCategory string  `json:"commodity_category"`
	Installments      int     `json:"installments"`
	PromoCode         string  `json:"promo_code,omitempty"`
	PurchaseDate      string  `json:"purchase_date,omitempty"`

	purchaseDate *time.Time
}

func (r *CreateRepaymentPlanRequest) Validate() error {
	if r.CustomerID == "" {
		return apperrors.MissingParamError("customer_id")
	}

	if math.IsNaN(r.CommodityCost) ||
		math.IsInf(r.CommodityCost, 0) ||
		r.CommodityCost <= 0 {
		return apperrors.InvalidCommodityCostError(r.CommodityCost)
	}

	if r.Installments < MinInstallments || r.Installments > MaxInstallments {
		return apperrors.InstallmentCountOutOfRangeError(r.Installments, MinInstallments, MaxInstallments)
	}

	if r.PurchaseDate == "" {
		return nil
	}

	parsed, err := time.Parse(purchaseDateLayout, r.PurchaseDate)
	if err != nil {
		return apperrors.InvalidRequestBodyError(err)
	}
	r.purchaseDate = &parsed
	return nil
}

func (r *CreateRepaymentPlanRequest) ToServiceInput() service.CreateRepaymentPlanInput {
	return service.CreateRepaymentPlanInput{
		CustomerID:        r.CustomerID,
		CommodityCost:     r.CommodityCost,
		CommodityCategory: r.CommodityCategory,
		Installments:      r.Installments,
		PromoCode:         r.PromoCode,
		PurchaseDate:      r.purchaseDate,
	}
}
