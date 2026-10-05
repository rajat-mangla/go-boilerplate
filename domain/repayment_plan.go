package domain

import (
	"fmt"
	"math"
	"time"
)

type ScheduledInstallment struct {
	DueDate time.Time
	Amount  Minor
}

type RepaymentPlan struct {
	CustomerID              string
	CommodityCost           Minor
	AppliedMarginHundredths int64
	ProfitAmount            Minor
	TotalPayable            Minor
	BaseInstallmentAmount   Minor
	Schedule                []ScheduledInstallment
	ContractSummary         string
}

type ComputeInput struct {
	CustomerID        string
	CommodityCost     float64
	CommodityCategory string
	Installments      int
	PromoCode         string
	PurchaseDate      time.Time
}

func Compute(input ComputeInput) (*RepaymentPlan, error) {
	discountPercent, err := ResolvePromoDiscount(input.PromoCode)
	if err != nil {
		return nil, err
	}

	// Compute the margin after applying any promotional discount
	marginHundredths := ApplyPromoDiscount(BaseMarginHundredths(input.CommodityCategory), discountPercent)

	cost := Minor(math.Round(input.CommodityCost * 100))
	profit := ApplyMargin(cost, marginHundredths)
	total := cost + profit

	installmentAmounts := SplitEven(total, input.Installments)
	dueDates := GenerateDueDates(input.PurchaseDate, input.Installments)

	schedule := make([]ScheduledInstallment, input.Installments)
	for i := range schedule {
		schedule[i] = ScheduledInstallment{DueDate: dueDates[i], Amount: installmentAmounts[i]}
	}

	return &RepaymentPlan{
		CustomerID:              input.CustomerID,
		CommodityCost:           cost,
		AppliedMarginHundredths: marginHundredths,
		ProfitAmount:            profit,
		TotalPayable:            total,
		BaseInstallmentAmount:   installmentAmounts[0],
		Schedule:                schedule,
		ContractSummary:         BuildContractSummary(cost, total, profit, installmentAmounts[0], input.Installments),
	}, nil
}

func BuildContractSummary(cost, total, profit, baseInstallment Minor, installments int) string {
	return fmt.Sprintf(
		"This Murabaha contract confirms our purchase of the commodity on your behalf for SAR %s. "+
			"We are selling it to you at a total price of SAR %s, which includes our profit of SAR %s. "+
			"This amount is to be paid in %d equal monthly installments of SAR %s.",
		cost.String(), total.String(), profit.String(), installments, baseInstallment.String(),
	)
}
