package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	apperrors "github.com/rajat-mangla/go-boilerplate/errors"
	"github.com/rajat-mangla/go-boilerplate/service"
)

func CreateRepaymentPlanV1(svc *service.RepaymentPlanService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateRepaymentPlanRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			WriteServiceError(w, apperrors.InvalidRequestBodyError(err))
			return
		}

		if err := req.Validate(); err != nil {
			writeRepaymentPlanError(w, err)
			return
		}

		result, err := svc.CreateRepaymentPlan(req.ToServiceInput())
		if err != nil {
			writeRepaymentPlanError(w, err)
			return
		}

		WriteJSON(w, http.StatusOK, toCreateRepaymentPlanResponse(result))
	}
}

func writeRepaymentPlanError(w http.ResponseWriter, err error) {
	var svcErr apperrors.ServiceError
	if errors.As(err, &svcErr) {
		WriteServiceError(w, svcErr)
		return
	}
	WriteError(w, http.StatusInternalServerError, "internal server error")
}
