package factory

import (
	"github.com/CarlosHe/go-policy-management/pkg/policy/evaluator/condition"
)

type IConditionFactory interface {
	CreateEvaluator() condition.Evaluator
	CreatePatternMatcher() condition.PatternMatcher
	CreateStringEvaluator() *condition.StringEvaluator
	CreateNumericEvaluator() *condition.NumericEvaluator
	CreateDateEvaluator() *condition.DateEvaluator
	CreateBoolEvaluator() *condition.BoolEvaluator
}

type DefaultConditionFactory struct{}

func NewConditionFactory() *DefaultConditionFactory {
	return &DefaultConditionFactory{}
}

func (f *DefaultConditionFactory) CreateEvaluator() condition.Evaluator {
	return condition.NewCompositeEvaluator()
}

func (f *DefaultConditionFactory) CreatePatternMatcher() condition.PatternMatcher {
	return condition.NewRegexPatternMatcher()
}

func (f *DefaultConditionFactory) CreateStringEvaluator() *condition.StringEvaluator {
	return condition.NewStringEvaluator(f.CreatePatternMatcher())
}

func (f *DefaultConditionFactory) CreateNumericEvaluator() *condition.NumericEvaluator {
	return condition.NewNumericEvaluator()
}

func (f *DefaultConditionFactory) CreateDateEvaluator() *condition.DateEvaluator {
	return condition.NewDateEvaluator()
}

func (f *DefaultConditionFactory) CreateBoolEvaluator() *condition.BoolEvaluator {
	return condition.NewBoolEvaluator()
}
