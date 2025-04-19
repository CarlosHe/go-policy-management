package condition

import (
	"github.com/CarlosHe/go-policy-management/pkg/policy"
)

type CompositeEvaluator struct {
	stringEvaluator  *StringEvaluator
	numericEvaluator *NumericEvaluator
	dateEvaluator    *DateEvaluator
	boolEvaluator    *BoolEvaluator
}

func NewCompositeEvaluator() *CompositeEvaluator {
	patternMatcher := NewRegexPatternMatcher()
	return &CompositeEvaluator{
		stringEvaluator:  NewStringEvaluator(patternMatcher),
		numericEvaluator: NewNumericEvaluator(),
		dateEvaluator:    NewDateEvaluator(),
		boolEvaluator:    NewBoolEvaluator(),
	}
}

func (e *CompositeEvaluator) Evaluate(condition policy.Condition, context map[string]interface{}) bool {
	key := string(condition.Key)
	contextValue, exists := context[key]
	if !exists {
		return false
	}
	switch condition.Operator {
	case policy.StringEquals:
		return e.stringEvaluator.Equals(contextValue, condition.Value)
	case policy.StringNotEquals:
		return e.stringEvaluator.NotEquals(contextValue, condition.Value)
	case policy.StringLike:
		return e.stringEvaluator.Like(contextValue, condition.Value)
	case policy.StringNotLike:
		return e.stringEvaluator.NotLike(contextValue, condition.Value)
	case policy.NumericEquals:
		return e.numericEvaluator.Equals(contextValue, condition.Value)
	case policy.NumericNotEquals:
		return e.numericEvaluator.NotEquals(contextValue, condition.Value)
	case policy.NumericLessThan:
		return e.numericEvaluator.LessThan(contextValue, condition.Value)
	case policy.NumericLessThanEquals:
		return e.numericEvaluator.LessThanEquals(contextValue, condition.Value)
	case policy.NumericGreaterThan:
		return e.numericEvaluator.GreaterThan(contextValue, condition.Value)
	case policy.NumericGreaterThanEquals:
		return e.numericEvaluator.GreaterThanEquals(contextValue, condition.Value)
	case policy.DateEquals:
		return e.dateEvaluator.Equals(contextValue, condition.Value)
	case policy.DateNotEquals:
		return e.dateEvaluator.NotEquals(contextValue, condition.Value)
	case policy.DateLessThan:
		return e.dateEvaluator.LessThan(contextValue, condition.Value)
	case policy.DateLessThanEquals:
		return e.dateEvaluator.LessThanEquals(contextValue, condition.Value)
	case policy.DateGreaterThan:
		return e.dateEvaluator.GreaterThan(contextValue, condition.Value)
	case policy.DateGreaterThanEquals:
		return e.dateEvaluator.GreaterThanEquals(contextValue, condition.Value)
	case policy.Bool:
		return e.boolEvaluator.Equals(contextValue, condition.Value)
	default:
		return false
	}
}
