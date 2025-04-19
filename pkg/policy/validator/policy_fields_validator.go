package validator

import (
	"fmt"

	"github.com/CarlosHe/go-policy-management/pkg/policy"
)

type PolicyFieldsValidator struct {
	RequireID     bool
	RequireName   bool
	MaxStatements int
}

func NewPolicyFieldsValidator() *PolicyFieldsValidator {
	return &PolicyFieldsValidator{
		RequireID:     true,
		RequireName:   true,
		MaxStatements: 100,
	}
}

func (v *PolicyFieldsValidator) ValidateFields(policy policy.Policy) []ValidationError {
	var errors []ValidationError

	if v.RequireID && policy.ID == "" {
		errors = append(errors, ValidationError{
			Field:   "ID",
			Message: "Policy ID is required",
		})
	}

	if v.RequireName && policy.Name == "" {
		errors = append(errors, ValidationError{
			Field:   "Name",
			Message: "Policy Name is required",
		})
	}

	if policy.Version == "" {
		errors = append(errors, ValidationError{
			Field:   "Version",
			Message: "Policy Version is required",
		})
	}

	if policy.CreatedAt.IsZero() {
		errors = append(errors, ValidationError{
			Field:   "CreatedAt",
			Message: "Creation date is required",
		})
	}

	if len(policy.Statements) == 0 {
		errors = append(errors, ValidationError{
			Field:   "Statements",
			Message: "Policy must have at least one statement",
		})
	}

	if v.MaxStatements > 0 && len(policy.Statements) > v.MaxStatements {
		errors = append(errors, ValidationError{
			Field:   "Statements",
			Message: fmt.Sprintf("Policy exceeds maximum number of statements (%d)", v.MaxStatements),
		})
	}

	return errors
}
