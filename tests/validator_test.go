package tests

import (
	"testing"

	"github.com/CarlosHe/go-policy-management/pkg/policy"
	"github.com/CarlosHe/go-policy-management/pkg/policy/factory"
	"github.com/CarlosHe/go-policy-management/pkg/policy/validator"
)

func TestValidatePolicy(t *testing.T) {
	// Arrange
	policyFactory := factory.NewPolicyFactory()
	policyValidator := validator.NewDefaultValidator()

	// Act - Create a valid policy
	validPolicy := policyFactory.CreatePolicy(
		"test-policy",
		"Test Policy",
		policyFactory.CreateStatement(
			"statement-1",
			policy.Allow,
			[]policy.Action{"read", "write"},
			[]policy.Resource{"resource:test:*"},
		),
	)

	// Assert - Verify that the valid policy does not generate errors
	errs := policyValidator.Validate(validPolicy)
	if len(errs) > 0 {
		t.Errorf("The valid policy should not generate errors: %v", errs)
	}

	// Act - Policy with empty ID
	invalidPolicy := policyFactory.CreatePolicy(
		"",
		"Test Policy",
		policyFactory.CreateStatement(
			"statement-1",
			policy.Allow,
			[]policy.Action{"read"},
			[]policy.Resource{"resource:test:*"},
		),
	)

	// Assert - Verify that the invalid policy generates errors
	errs = policyValidator.Validate(invalidPolicy)
	if len(errs) == 0 {
		t.Errorf("Should return an error for policy with empty ID")
	}

	// Verify that the error is for the ID field
	found := false
	for _, err := range errs {
		if err.Field == "ID" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Should include error for the ID field, but received: %v", errs)
	}
}

func TestValidateStatement(t *testing.T) {
	// Arrange
	policyFactory := factory.NewPolicyFactory()
	statementValidator := validator.NewStatementValidator()

	// Act - Create a valid statement
	validStatement := policyFactory.CreateStatement(
		"statement-1",
		policy.Allow,
		[]policy.Action{"read", "write"},
		[]policy.Resource{"resource:test:*"},
	)

	// Assert - Verify that the valid statement does not generate errors
	errs := statementValidator.ValidateStatement(validStatement, 0)
	if len(errs) > 0 {
		t.Errorf("The valid statement should not generate errors: %v", errs)
	}

	// Act - Statement with invalid effect
	invalidStatement := policy.Statement{
		ID:        "statement-1",
		Effect:    "INVALID_EFFECT",
		Actions:   []policy.Action{"read"},
		Resources: []policy.Resource{"resource:test:*"},
	}

	// Assert - Verify that the invalid statement generates errors
	errs = statementValidator.ValidateStatement(invalidStatement, 0)
	if len(errs) == 0 {
		t.Errorf("Should return an error for statement with invalid effect")
	}

	// Verify that the error is for the Effect field
	found := false
	for _, err := range errs {
		if err.Field == "Statements[0].Effect" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Should include error for the Effect field, but received: %v", errs)
	}

	// Act - Statement without actions
	noActionsStatement := policy.Statement{
		ID:        "statement-1",
		Effect:    policy.Allow,
		Actions:   []policy.Action{},
		Resources: []policy.Resource{"resource:test:*"},
	}

	// Assert - Verify that the statement without actions generates errors
	errs = statementValidator.ValidateStatement(noActionsStatement, 0)
	if len(errs) == 0 {
		t.Errorf("Should return an error for statement without actions")
	}

	// Verify that the error is for the Actions field
	found = false
	for _, err := range errs {
		if err.Field == "Statements[0].Actions" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Should include error for the Actions field, but received: %v", errs)
	}
}

func TestValidateCondition(t *testing.T) {
	// Arrange
	conditionValidator := validator.NewConditionValidator()

	// Act - Create a valid condition
	validCondition := policy.Condition{
		Operator: "StringEquals",
		Key:      "user.role",
		Value:    "admin",
	}

	// Assert - Verify that the valid condition does not generate errors
	errs := conditionValidator.ValidateCondition(validCondition, 0, 0)
	if len(errs) > 0 {
		t.Errorf("The valid condition should not generate errors: %v", errs)
	}

	// Act - Condition with empty operator
	invalidCondition := policy.Condition{
		Operator: "",
		Key:      "user.role",
		Value:    "admin",
	}

	// Assert - Verify that the invalid condition generates errors
	errs = conditionValidator.ValidateCondition(invalidCondition, 0, 0)
	if len(errs) == 0 {
		t.Errorf("Should return an error for condition with empty operator")
	}

	// Verify that the error is for the Operator field
	found := false
	for _, err := range errs {
		if err.Field == "Statements[0].Conditions[0].Operator" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Should include error for the Operator field, but received: %v", errs)
	}

	// Act - Condition with empty key
	emptyKeyCondition := policy.Condition{
		Operator: "StringEquals",
		Key:      "",
		Value:    "admin",
	}

	// Assert - Verify that the condition with empty key generates errors
	errs = conditionValidator.ValidateCondition(emptyKeyCondition, 0, 0)
	if len(errs) == 0 {
		t.Errorf("Should return an error for condition with empty key")
	}

	// Verify that the error is for the Key field
	found = false
	for _, err := range errs {
		if err.Field == "Statements[0].Conditions[0].Key" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Should include error for the Key field, but received: %v", errs)
	}
}
