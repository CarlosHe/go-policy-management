package main

import (
	"fmt"

	"github.com/CarlosHe/go-policy-management/pkg/policy"
	"github.com/CarlosHe/go-policy-management/pkg/policy/evaluator"
	"github.com/CarlosHe/go-policy-management/pkg/policy/factory"
)

func main() {
	policyFactory := factory.NewPolicyFactory()
	validatorFactory := factory.NewValidatorFactory()
	evaluatorFactory := factory.NewEvaluatorFactory()

	readOnlyPolicy := policyFactory.CreatePolicy(
		"p-1234",
		"ReadOnlyAccess",
		policyFactory.CreateStatement(
			"s-1",
			policy.Allow,
			[]policy.Action{"read", "list", "get"},
			[]policy.Resource{"resource:*"},
		),
		policyFactory.CreateStatement(
			"s-2",
			policy.Deny,
			[]policy.Action{"write", "update", "delete", "create"},
			[]policy.Resource{"resource:*"},
		),
	)

	validator := validatorFactory.CreatePolicyValidator()
	errors := validator.Validate(readOnlyPolicy)
	if len(errors) > 0 {
		fmt.Println("Validation errors:")
		for _, err := range errors {
			fmt.Printf("- %s: %s\n", err.Field, err.Message)
		}
		return
	}

	eval := evaluatorFactory.CreatePolicyEvaluator(readOnlyPolicy)

	allowRequest := evaluator.Request{
		Principal: "user-123",
		Action:    "read",
		Resource:  "resource:doc1",
		Context: map[string]interface{}{
			"time": "2023-05-01T10:00:00Z",
		},
	}

	denyRequest := evaluator.Request{
		Principal: "user-123",
		Action:    "write",
		Resource:  "resource:doc1",
		Context: map[string]interface{}{
			"time": "2023-05-01T10:00:00Z",
		},
	}

	allowResult := eval.Evaluate(allowRequest)
	fmt.Printf("Allow request result: allowed=%v, reason=%s\n",
		allowResult.Allowed, allowResult.Reason)

	denyResult := eval.Evaluate(denyRequest)
	fmt.Printf("Deny request result: allowed=%v, reason=%s\n",
		denyResult.Allowed, denyResult.Reason)
}
