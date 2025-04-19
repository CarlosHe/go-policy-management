package validator

import (
	"github.com/CarlosHe/go-policy-management/pkg/policy"
)

type DefaultValidator struct {
	policyFieldsValidator IPolicyFieldsValidator
	policyValidator       *PolicyValidator
}

func NewDefaultValidator() *DefaultValidator {
	statementValidator := NewStatementValidator()
	conditionValidator := NewConditionValidator()

	return &DefaultValidator{
		policyFieldsValidator: NewPolicyFieldsValidator(),
		policyValidator:       NewPolicyValidator(statementValidator, conditionValidator),
	}
}

func (v *DefaultValidator) Validate(policy policy.Policy) []ValidationError {
	errors := v.policyFieldsValidator.ValidateFields(policy)

	evalErrors := v.policyValidator.ValidatePolicy(policy)
	errors = append(errors, evalErrors...)

	return errors
}

func (v *DefaultValidator) ValidateStatement(statement policy.Statement, index int) []ValidationError {
	return v.policyValidator.ValidateStatement(statement, index)
}

func (v *DefaultValidator) ValidateCondition(condition policy.Condition, stmIndex, condIndex int) []ValidationError {
	return v.policyValidator.ValidateCondition(condition, stmIndex, condIndex)
}
