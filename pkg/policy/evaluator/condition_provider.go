package evaluator

import (
	"github.com/CarlosHe/go-policy-management/pkg/policy/evaluator/condition"
)

type IConditionProvider interface {
	GetEvaluator() condition.Evaluator
	GetPatternMatcher() condition.PatternMatcher
}
