package evaluator

import (
	"github.com/CarlosHe/go-policy-management/pkg/policy"
)

type DefaultPolicyEvaluator struct {
	policies          []policy.Policy
	conditionProvider IConditionProvider
	policyMatcher     *PolicyMatcher
}

func NewDefaultEvaluator(conditionProvider IConditionProvider, policies ...policy.Policy) *DefaultPolicyEvaluator {
	return &DefaultPolicyEvaluator{
		policies:          policies,
		conditionProvider: conditionProvider,
		policyMatcher:     NewPolicyMatcher(conditionProvider),
	}
}

func (e *DefaultPolicyEvaluator) AddPolicy(policy policy.Policy) {
	e.policies = append(e.policies, policy)
}

func (e *DefaultPolicyEvaluator) Evaluate(req Request) Result {
	return e.policyMatcher.MatchPolicy(req, e.policies)
}
