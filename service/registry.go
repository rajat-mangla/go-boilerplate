package service

import (
	"github.com/rajat-mangla/go-boilerplate/config"
)

type Registry struct {
	RepaymentPlan *RepaymentPlanService
}

func NewRegistry(_ config.Config) *Registry {
	return &Registry{
		RepaymentPlan: NewRepaymentPlanService(),
	}
}
