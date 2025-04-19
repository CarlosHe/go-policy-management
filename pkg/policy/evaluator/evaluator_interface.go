package evaluator

import (
	"github.com/CarlosHe/go-policy-management/pkg/policy"
)

type IPolicyEvaluator interface {
	Evaluate(req Request) Result
	AddPolicy(policy policy.Policy)
}
