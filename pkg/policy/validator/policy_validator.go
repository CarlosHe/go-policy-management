package validator

import (
	"github.com/CarlosHe/go-policy-management/pkg/policy"
)

type PolicyValidator struct {
	statementValidator IStatementValidator
	conditionValidator IConditionValidator
}

func NewPolicyValidator(
	statementValidator IStatementValidator,
	conditionValidator IConditionValidator,
) *PolicyValidator {
	return &PolicyValidator{
		statementValidator: statementValidator,
		conditionValidator: conditionValidator,
	}
}

func (v *PolicyValidator) ValidatePolicy(policy policy.Policy) []ValidationError {
	var errors []ValidationError

	for i, statement := range policy.Statements {
		stmErrors := v.ValidateStatement(statement, i)
		errors = append(errors, stmErrors...)
	}

	return errors
}

func (v *PolicyValidator) ValidateStatement(statement policy.Statement, index int) []ValidationError {
	errors := v.statementValidator.ValidateStatement(statement, index)

	for j, condition := range statement.Conditions {
		condErrors := v.ValidateCondition(condition, index, j)
		errors = append(errors, condErrors...)
	}

	return errors
}

func (v *PolicyValidator) ValidateCondition(condition policy.Condition, stmIndex, condIndex int) []ValidationError {
	return v.conditionValidator.ValidateCondition(condition, stmIndex, condIndex)
}
