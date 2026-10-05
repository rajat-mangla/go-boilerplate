package handlers

import "github.com/rajat-mangla/go-boilerplate/service"

type InstallmentResponse struct {
	DueDate string  `json:"due_date"`
	Amount  float64 `json:"amount"`
}

type CreateRepaymentPlanResponse struct {
	RepaymentPlanID         string                `json:"repayment_plan_id"`
	CustomerID              string                `json:"customer_id"`
	CommodityCost           float64               `json:"commodity_cost"`
	AppliedMarginPercentage float64               `json:"applied_margin_percentage"`
	ProfitAmount            float64               `json:"profit_amount"`
	TotalPayable            float64               `json:"total_payable"`
	BaseInstallmentAmount   float64               `json:"base_installment_amount"`
	RepaymentSchedule       []InstallmentResponse `json:"repayment_schedule"`
	ContractSummary         string                `json:"contract_summary"`
}

// toCreateRepaymentPlanResponse maps a service result to the response DTO.
func toCreateRepaymentPlanResponse(result *service.RepaymentPlanResult) CreateRepaymentPlanResponse {
	plan := result.Plan

	schedule := make([]InstallmentResponse, len(plan.Schedule))
	for i, installment := range plan.Schedule {
		schedule[i] = InstallmentResponse{
			DueDate: installment.DueDate.Format(purchaseDateLayout),
			Amount:  installment.Amount.SAR(),
		}
	}

	return CreateRepaymentPlanResponse{
		RepaymentPlanID:         result.ID,
		CustomerID:              plan.CustomerID,
		CommodityCost:           plan.CommodityCost.SAR(),
		AppliedMarginPercentage: float64(plan.AppliedMarginHundredths) / 100,
		ProfitAmount:            plan.ProfitAmount.SAR(),
		TotalPayable:            plan.TotalPayable.SAR(),
		BaseInstallmentAmount:   plan.BaseInstallmentAmount.SAR(),
		RepaymentSchedule:       schedule,
		ContractSummary:         plan.ContractSummary,
	}
}
