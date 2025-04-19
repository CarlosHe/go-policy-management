package factory

import (
	"github.com/CarlosHe/go-policy-management/pkg/policy/evaluator/condition"
)

type ConditionFactoryAdapter struct {
	factory IConditionFactory
}

func NewConditionFactoryAdapter(factory IConditionFactory) *ConditionFactoryAdapter {
	return &ConditionFactoryAdapter{
		factory: factory,
	}
}

func (a *ConditionFactoryAdapter) GetEvaluator() condition.Evaluator {
	return a.factory.CreateEvaluator()
}

func (a *ConditionFactoryAdapter) GetPatternMatcher() condition.PatternMatcher {
	return a.factory.CreatePatternMatcher()
}
