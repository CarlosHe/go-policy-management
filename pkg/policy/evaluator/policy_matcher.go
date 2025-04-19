package evaluator

import (
	"fmt"
	"time"

	"github.com/CarlosHe/go-policy-management/pkg/policy"
)

type PolicyMatcher struct {
	conditionProvider IConditionProvider
}

func NewPolicyMatcher(conditionProvider IConditionProvider) *PolicyMatcher {
	return &PolicyMatcher{
		conditionProvider: conditionProvider,
	}
}

func (m *PolicyMatcher) MatchPolicy(req Request, policies []policy.Policy) Result {
	result := Result{
		Allowed:     false,
		EvaluatedAt: time.Now(),
	}
	if len(policies) == 0 {
		result.Reason = "No policies defined"
		return result
	}
	for _, p := range policies {
		for _, statement := range p.Statements {
			if m.matchStatement(req, statement, p.ID) {
				if statement.Effect == policy.Allow {
					result.Allowed = true
					result.Reason = fmt.Sprintf("Allowed by policy %s, statement %s", p.ID, statement.ID)
					result.MatchedRules = append(result.MatchedRules, statement.ID)
				} else if statement.Effect == policy.Deny {
					result.Allowed = false
					result.Reason = fmt.Sprintf("Denied by policy %s, statement %s", p.ID, statement.ID)
					result.MatchedRules = append(result.MatchedRules, statement.ID)
					return result
				}
			}
		}
	}

	return result
}

func (m *PolicyMatcher) matchStatement(req Request, statement policy.Statement, policyID string) bool {
	actionMatched := false
	for _, action := range statement.Actions {
		if m.conditionProvider.GetPatternMatcher().MatchesPattern(string(req.Action), string(action)) {
			actionMatched = true
			break
		}
	}
	if !actionMatched {
		return false
	}

	resourceMatched := false
	for _, resource := range statement.Resources {
		if m.conditionProvider.GetPatternMatcher().MatchesPattern(string(req.Resource), string(resource)) {
			resourceMatched = true
			break
		}
	}
	if !resourceMatched {
		return false
	}
	for _, condition := range statement.Conditions {
		if !m.conditionProvider.GetEvaluator().Evaluate(condition, req.Context) {
			return false
		}
	}
	return true
}
