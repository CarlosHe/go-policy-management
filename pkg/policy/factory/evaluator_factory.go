package factory

import (
	"github.com/CarlosHe/go-policy-management/pkg/policy"
	"github.com/CarlosHe/go-policy-management/pkg/policy/evaluator"
)

type IEvaluatorFactory interface {
	CreatePolicyEvaluator(policies ...policy.Policy) evaluator.IPolicyEvaluator
}

type DefaultEvaluatorFactory struct {
	conditionFactory IConditionFactory
}

func NewEvaluatorFactory() *DefaultEvaluatorFactory {
	return &DefaultEvaluatorFactory{
		conditionFactory: NewConditionFactory(),
	}
}

func (f *DefaultEvaluatorFactory) CreatePolicyEvaluator(policies ...policy.Policy) evaluator.IPolicyEvaluator {
	adapter := NewConditionFactoryAdapter(f.conditionFactory)
	return evaluator.NewDefaultEvaluator(adapter, policies...)
}
