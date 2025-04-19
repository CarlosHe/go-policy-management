package validator

import (
	"fmt"

	"github.com/CarlosHe/go-policy-management/pkg/policy"
)

type StatementValidator struct {
	MaxActionsPerStm    int
	MaxResourcesPerStm  int
	MaxConditionsPerStm int
}

func NewStatementValidator() *StatementValidator {
	return &StatementValidator{
		MaxActionsPerStm:    50,
		MaxResourcesPerStm:  100,
		MaxConditionsPerStm: 20,
	}
}

func (v *StatementValidator) ValidateStatement(statement policy.Statement, index int) []ValidationError {
	var errors []ValidationError
	fieldPrefix := fmt.Sprintf("Statements[%d].", index)

	if statement.Effect != policy.Allow && statement.Effect != policy.Deny {
		errors = append(errors, ValidationError{
			Field:   fieldPrefix + "Effect",
			Message: "Effect must be either Allow or Deny",
		})
	}

	if len(statement.Actions) == 0 {
		errors = append(errors, ValidationError{
			Field:   fieldPrefix + "Actions",
			Message: "Statement must have at least one action",
		})
	}

	if v.MaxActionsPerStm > 0 && len(statement.Actions) > v.MaxActionsPerStm {
		errors = append(errors, ValidationError{
			Field:   fieldPrefix + "Actions",
			Message: fmt.Sprintf("Statement exceeds maximum number of actions (%d)", v.MaxActionsPerStm),
		})
	}

	if len(statement.Resources) == 0 {
		errors = append(errors, ValidationError{
			Field:   fieldPrefix + "Resources",
			Message: "Statement must have at least one resource",
		})
	}

	if v.MaxResourcesPerStm > 0 && len(statement.Resources) > v.MaxResourcesPerStm {
		errors = append(errors, ValidationError{
			Field:   fieldPrefix + "Resources",
			Message: fmt.Sprintf("Statement exceeds maximum number of resources (%d)", v.MaxResourcesPerStm),
		})
	}

	if v.MaxConditionsPerStm > 0 && len(statement.Conditions) > v.MaxConditionsPerStm {
		errors = append(errors, ValidationError{
			Field:   fieldPrefix + "Conditions",
			Message: fmt.Sprintf("Statement exceeds maximum number of conditions (%d)", v.MaxConditionsPerStm),
		})
	}

	return errors
}
