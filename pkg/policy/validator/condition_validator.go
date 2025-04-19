package validator

import (
	"fmt"

	"github.com/CarlosHe/go-policy-management/pkg/policy"
)

type ConditionValidator struct{}

func NewConditionValidator() *ConditionValidator {
	return &ConditionValidator{}
}

func (v *ConditionValidator) ValidateCondition(condition policy.Condition, stmIndex, condIndex int) []ValidationError {
	var errors []ValidationError
	fieldPrefix := fmt.Sprintf("Statements[%d].Conditions[%d].", stmIndex, condIndex)

	if condition.Operator == "" {
		errors = append(errors, ValidationError{
			Field:   fieldPrefix + "Operator",
			Message: "Condition operator is required",
		})
	}

	if condition.Key == "" {
		errors = append(errors, ValidationError{
			Field:   fieldPrefix + "Key",
			Message: "Condition key is required",
		})
	}

	if condition.Value == nil {
		errors = append(errors, ValidationError{
			Field:   fieldPrefix + "Value",
			Message: "Condition value cannot be nil",
		})
	}

	return errors
}
