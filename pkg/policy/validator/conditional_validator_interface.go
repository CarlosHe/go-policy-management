package validator

import "github.com/CarlosHe/go-policy-management/pkg/policy"

type IConditionValidator interface {
	ValidateCondition(condition policy.Condition, stmIndex, condIndex int) []ValidationError
}
