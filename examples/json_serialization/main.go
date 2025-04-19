package main

import (
	"fmt"
	"time"

	"github.com/CarlosHe/go-policy-management/pkg/policy"
	"github.com/CarlosHe/go-policy-management/pkg/policy/evaluator"
	"github.com/CarlosHe/go-policy-management/pkg/policy/factory"
)

func main() {
	fmt.Println("=== Policy JSON Serialization Example ===")

	policyFactory := factory.NewPolicyFactory()

	readPolicy := policyFactory.CreatePolicy(
		"p-read-001",
		"ReadOnlyAccess",
		policyFactory.CreateStatement(
			"s-1",
			policy.Allow,
			[]policy.Action{"read", "list"},
			[]policy.Resource{"resource:document:*"},
		),
	)

	adminStatement := policyFactory.CreateStatement(
		"s-1",
		policy.Allow,
		[]policy.Action{"*"},
		[]policy.Resource{"*"},
	)

	adminStatement.Conditions = []policy.Condition{
		{
			Operator: policy.StringEquals,
			Key:      "user.role",
			Value:    "admin",
		},
	}

	adminPolicy := policyFactory.CreatePolicy(
		"p-admin-001",
		"AdminAccess",
		adminStatement,
	)

	fmt.Println("\n=== Policy to JSON ===")
	readPolicyJSON, err := readPolicy.ToJSONIndent()
	if err != nil {
		fmt.Printf("Error converting policy to JSON: %v\n", err)
		return
	}
	fmt.Println(readPolicyJSON)

	fmt.Println("\n=== JSON to Policy ===")
	jsonStr := `{
		"version": "2023-01-01",
		"id": "p-custom-001",
		"name": "CustomPolicy",
		"description": "Policy created from JSON",
		"statements": [
			{
				"id": "s-1",
				"effect": "Allow",
				"actions": ["read", "write"],
				"resources": ["resource:custom:*"]
			}
		],
		"created_at": "2023-05-10T15:04:05Z"
	}`

	customPolicy, err := policy.FromJSON(jsonStr)
	if err != nil {
		fmt.Printf("Error converting JSON to policy: %v\n", err)
		return
	}

	fmt.Printf("Policy loaded from JSON:\n")
	fmt.Printf("  ID: %s\n", customPolicy.ID)
	fmt.Printf("  Name: %s\n", customPolicy.Name)
	fmt.Printf("  Description: %s\n", customPolicy.Description)
	fmt.Printf("  Statements: %d\n", len(customPolicy.Statements))
	fmt.Printf("  Created at: %s\n", customPolicy.CreatedAt.Format(time.RFC3339))

	fmt.Println("\n=== Policy List to JSON ===")
	policies := []policy.Policy{readPolicy, adminPolicy, customPolicy}
	policiesJSON, err := policy.ToJSONListIndent(policies)
	if err != nil {
		fmt.Printf("Error converting policy list to JSON: %v\n", err)
		return
	}
	fmt.Println(policiesJSON)

	fmt.Println("\n=== JSON to Policy List ===")
	jsonListStr := `{
		"policies": [
			{
				"version": "2023-01-01",
				"id": "p-json-001",
				"name": "JSONPolicy1",
				"statements": [
					{
						"id": "s-1",
						"effect": "Allow",
						"actions": ["read"],
						"resources": ["resource:json:*"]
					}
				],
				"created_at": "2023-05-11T10:00:00Z"
			},
			{
				"version": "2023-01-01",
				"id": "p-json-002",
				"name": "JSONPolicy2",
				"statements": [
					{
						"id": "s-1",
						"effect": "Deny",
						"actions": ["delete"],
						"resources": ["resource:json:sensitive/*"]
					}
				],
				"created_at": "2023-05-11T11:00:00Z"
			}
		]
	}`

	jsonPolicies, err := policy.FromJSONList(jsonListStr)
	if err != nil {
		fmt.Printf("Error converting JSON to policy list: %v\n", err)
		return
	}

	fmt.Printf("Policies loaded from JSON: %d\n", len(jsonPolicies))
	for i, p := range jsonPolicies {
		fmt.Printf("  Policy %d: %s (%s)\n", i+1, p.Name, p.ID)
	}

	fmt.Println("\n=== Demonstration: Policy Loading Workflow ===")
	fmt.Println("1. Load policies from a JSON file (simulated)")
	fmt.Println("2. Validate each policy before using")
	fmt.Println("3. Create an evaluator with the loaded policies")

	jsonContent := jsonListStr

	loadedPolicies, err := policy.FromJSONList(jsonContent)
	if err != nil {
		fmt.Printf("Error loading policies: %v\n", err)
		return
	}

	validatorFactory := factory.NewValidatorFactory()
	validator := validatorFactory.CreatePolicyValidator()

	var validPolicies []policy.Policy
	for i, p := range loadedPolicies {
		errors := validator.Validate(p)
		if len(errors) > 0 {
			fmt.Printf("Policy %d (%s) invalid:\n", i+1, p.ID)
			for _, err := range errors {
				fmt.Printf("  - %s: %s\n", err.Field, err.Message)
			}
			continue
		}
		validPolicies = append(validPolicies, p)
	}

	fmt.Printf("Valid policies: %d/%d\n", len(validPolicies), len(loadedPolicies))

	evaluatorFactory := factory.NewEvaluatorFactory()
	eval := evaluatorFactory.CreatePolicyEvaluator(validPolicies...)

	fmt.Println("Evaluator successfully created with policies loaded from JSON!")

	request := evaluator.Request{
		Principal: "user-123",
		Action:    "read",
		Resource:  "resource:json:document1",
		Context: map[string]interface{}{
			"user.role": "standard",
		},
	}

	result := eval.Evaluate(request)
	fmt.Printf("Access evaluation: %v (%s)\n", result.Allowed, result.Reason)
}
