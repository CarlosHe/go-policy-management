package tests

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/CarlosHe/go-policy-management/pkg/policy"
	"github.com/CarlosHe/go-policy-management/pkg/policy/factory"
)

func TestPolicyMarshalJSON(t *testing.T) {
	// Arrange
	testTime := time.Date(2023, 5, 1, 10, 0, 0, 0, time.UTC)

	// Create a policy for testing
	testPolicy := policy.Policy{
		Version:     policy.PolicyVersion,
		ID:          "test-policy",
		Name:        "Test Policy",
		Description: "Policy for testing",
		Statements: []policy.Statement{
			{
				ID:     "statement-1",
				Effect: policy.Allow,
				Actions: []policy.Action{
					"read",
					"list",
				},
				Resources: []policy.Resource{
					"resource:test:*",
				},
			},
		},
		CreatedAt: testTime,
		UpdatedAt: testTime.Add(24 * time.Hour),
	}

	// Act
	bytes, err := json.Marshal(testPolicy)

	// Assert
	if err != nil {
		t.Fatalf("Failed to serialize policy: %v", err)
	}

	// Verify if the JSON contains the expected fields
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(bytes, &jsonMap); err != nil {
		t.Fatalf("Failed to parse resulting JSON: %v", err)
	}

	// Verify main fields
	assertJSONField(t, jsonMap, "id", "test-policy")
	assertJSONField(t, jsonMap, "name", "Test Policy")
	assertJSONField(t, jsonMap, "description", "Policy for testing")
	assertJSONField(t, jsonMap, "version", policy.PolicyVersion)
	assertJSONField(t, jsonMap, "created_at", "2023-05-01T10:00:00Z")
	assertJSONField(t, jsonMap, "updated_at", "2023-05-02T10:00:00Z")

	// Verify statements
	statements, ok := jsonMap["statements"].([]interface{})
	if !ok || len(statements) != 1 {
		t.Fatalf("Invalid statements structure or incorrect size")
	}

	statement := statements[0].(map[string]interface{})
	assertJSONField(t, statement, "id", "statement-1")
	assertJSONField(t, statement, "effect", "Allow")

	// Verify actions
	actions, ok := statement["actions"].([]interface{})
	if !ok || len(actions) != 2 {
		t.Fatalf("Invalid actions structure or incorrect size")
	}

	if actions[0].(string) != "read" || actions[1].(string) != "list" {
		t.Errorf("Incorrect actions in JSON: %v", actions)
	}
}

func TestPolicyUnmarshalJSON(t *testing.T) {
	// Arrange
	jsonStr := `{
		"version": "2023-01-01",
		"id": "test-policy",
		"name": "Test Policy",
		"description": "Policy for testing",
		"statements": [
			{
				"id": "statement-1",
				"effect": "Allow",
				"actions": ["read", "list"],
				"resources": ["resource:test:*"]
			}
		],
		"created_at": "2023-05-01T10:00:00Z",
		"updated_at": "2023-05-02T10:00:00Z"
	}`

	// Act
	var result policy.Policy
	err := json.Unmarshal([]byte(jsonStr), &result)

	// Assert
	if err != nil {
		t.Fatalf("Failed to deserialize policy: %v", err)
	}

	// Verify main fields
	if result.ID != "test-policy" {
		t.Errorf("Incorrect ID: expected 'test-policy', got '%s'", result.ID)
	}

	if result.Name != "Test Policy" {
		t.Errorf("Incorrect Name: expected 'Test Policy', got '%s'", result.Name)
	}

	if result.Description != "Policy for testing" {
		t.Errorf("Incorrect Description: expected 'Policy for testing', got '%s'", result.Description)
	}

	if result.Version != policy.PolicyVersion {
		t.Errorf("Incorrect Version: expected '%s', got '%s'", policy.PolicyVersion, result.Version)
	}

	// Verify dates
	expectedCreatedAt := time.Date(2023, 5, 1, 10, 0, 0, 0, time.UTC)
	if !result.CreatedAt.Equal(expectedCreatedAt) {
		t.Errorf("Incorrect CreatedAt: expected '%v', got '%v'", expectedCreatedAt, result.CreatedAt)
	}

	expectedUpdatedAt := time.Date(2023, 5, 2, 10, 0, 0, 0, time.UTC)
	if !result.UpdatedAt.Equal(expectedUpdatedAt) {
		t.Errorf("Incorrect UpdatedAt: expected '%v', got '%v'", expectedUpdatedAt, result.UpdatedAt)
	}

	// Verify statements
	if len(result.Statements) != 1 {
		t.Fatalf("Incorrect number of statements: expected 1, got %d", len(result.Statements))
	}

	statement := result.Statements[0]
	if statement.ID != "statement-1" {
		t.Errorf("Incorrect Statement ID: expected 'statement-1', got '%s'", statement.ID)
	}

	if statement.Effect != policy.Allow {
		t.Errorf("Incorrect Statement Effect: expected '%s', got '%s'", policy.Allow, statement.Effect)
	}

	// Verify actions
	if len(statement.Actions) != 2 {
		t.Fatalf("Incorrect number of actions: expected 2, got %d", len(statement.Actions))
	}

	if statement.Actions[0] != "read" || statement.Actions[1] != "list" {
		t.Errorf("Incorrect Actions: expected ['read','list'], got %v", statement.Actions)
	}
}

func TestFromJSON(t *testing.T) {
	// Arrange
	jsonStr := `{
		"version": "2023-01-01",
		"id": "test-policy",
		"name": "Test Policy",
		"statements": [
			{
				"id": "statement-1",
				"effect": "Allow",
				"actions": ["read"],
				"resources": ["resource:test:*"]
			}
		],
		"created_at": "2023-05-01T10:00:00Z"
	}`

	// Act
	result, err := policy.FromJSON(jsonStr)

	// Assert
	if err != nil {
		t.Fatalf("Failed to convert JSON to policy: %v", err)
	}

	if result.ID != "test-policy" {
		t.Errorf("Incorrect ID: expected 'test-policy', got '%s'", result.ID)
	}

	if len(result.Statements) != 1 {
		t.Errorf("Incorrect number of statements: expected 1, got %d", len(result.Statements))
	}
}

func TestToJSON(t *testing.T) {
	// Arrange
	policyFactory := factory.NewPolicyFactory()
	testPolicy := policyFactory.CreatePolicy(
		"test-policy",
		"Test Policy",
		policyFactory.CreateStatement(
			"statement-1",
			policy.Allow,
			[]policy.Action{"read"},
			[]policy.Resource{"resource:test:*"},
		),
	)

	// Act
	jsonStr, err := testPolicy.ToJSON()

	// Assert
	if err != nil {
		t.Fatalf("Failed to convert policy to JSON: %v", err)
	}

	// Verify if the JSON can be parsed back to a policy
	result, err := policy.FromJSON(jsonStr)
	if err != nil {
		t.Fatalf("Failed to convert JSON back to policy: %v", err)
	}

	if result.ID != "test-policy" {
		t.Errorf("Incorrect ID after round-trip conversion: expected 'test-policy', got '%s'", result.ID)
	}
}

func TestFromJSONList(t *testing.T) {
	// Arrange
	jsonStr := `{
		"policies": [
			{
				"version": "2023-01-01",
				"id": "policy-1",
				"name": "Policy 1",
				"statements": [
					{
						"id": "statement-1",
						"effect": "Allow",
						"actions": ["read"],
						"resources": ["resource:test:*"]
					}
				],
				"created_at": "2023-05-01T10:00:00Z"
			},
			{
				"version": "2023-01-01",
				"id": "policy-2",
				"name": "Policy 2",
				"statements": [
					{
						"id": "statement-1",
						"effect": "Deny",
						"actions": ["delete"],
						"resources": ["resource:test:sensitive/*"]
					}
				],
				"created_at": "2023-05-01T10:00:00Z"
			}
		]
	}`

	// Act
	policies, err := policy.FromJSONList(jsonStr)

	// Assert
	if err != nil {
		t.Fatalf("Failed to convert JSON to policy list: %v", err)
	}

	if len(policies) != 2 {
		t.Fatalf("Incorrect number of policies: expected 2, got %d", len(policies))
	}

	if policies[0].ID != "policy-1" {
		t.Errorf("Incorrect ID for first policy: expected 'policy-1', got '%s'", policies[0].ID)
	}

	if policies[1].ID != "policy-2" {
		t.Errorf("Incorrect ID for second policy: expected 'policy-2', got '%s'", policies[1].ID)
	}
}

func TestToJSONList(t *testing.T) {
	// Arrange
	policyFactory := factory.NewPolicyFactory()
	policies := []policy.Policy{
		policyFactory.CreatePolicy(
			"policy-1",
			"Policy 1",
			policyFactory.CreateStatement(
				"statement-1",
				policy.Allow,
				[]policy.Action{"read"},
				[]policy.Resource{"resource:test:*"},
			),
		),
		policyFactory.CreatePolicy(
			"policy-2",
			"Policy 2",
			policyFactory.CreateStatement(
				"statement-1",
				policy.Deny,
				[]policy.Action{"delete"},
				[]policy.Resource{"resource:test:sensitive/*"},
			),
		),
	}

	// Act
	jsonStr, err := policy.ToJSONList(policies)

	// Assert
	if err != nil {
		t.Fatalf("Failed to convert policy list to JSON: %v", err)
	}

	// Verify if the JSON can be parsed back to a policy list
	result, err := policy.FromJSONList(jsonStr)
	if err != nil {
		t.Fatalf("Failed to convert JSON back to policy list: %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("Incorrect number of policies after round-trip conversion: expected 2, got %d", len(result))
	}

	if result[0].ID != "policy-1" {
		t.Errorf("Incorrect ID for first policy after round-trip conversion: expected 'policy-1', got '%s'", result[0].ID)
	}
}

func TestInvalidJSON(t *testing.T) {
	// Arrange - Invalid JSON (key without quotes)
	invalidJSON := `{
		version: "2023-01-01",
		"id": "test-policy",
		"created_at": "2023-05-01T10:00:00Z"
	}`

	// Act
	_, err := policy.FromJSON(invalidJSON)

	// Assert
	if err == nil {
		t.Errorf("Expected an error when parsing invalid JSON, but none was returned")
	}
}

func TestMissingCreatedAt(t *testing.T) {
	// Arrange - JSON without created_at field
	invalidJSON := `{
		"version": "2023-01-01",
		"id": "test-policy",
		"name": "Test Policy",
		"statements": []
	}`

	// Act
	_, err := policy.FromJSON(invalidJSON)

	// Assert
	if err == nil {
		t.Errorf("Expected an error due to missing created_at field, but none was returned")
	}
}

func TestInvalidCreatedAtFormat(t *testing.T) {
	// Arrange - Invalid format for created_at
	invalidJSON := `{
		"version": "2023-01-01",
		"id": "test-policy",
		"name": "Test Policy",
		"statements": [],
		"created_at": "2023/05/01 10:00:00"
	}`

	// Act
	_, err := policy.FromJSON(invalidJSON)

	// Assert
	if err == nil {
		t.Errorf("Expected an error due to invalid created_at format, but none was returned")
	}
}

// Helper function to verify fields in JSON
func assertJSONField(t *testing.T, jsonObj map[string]interface{}, field string, expected interface{}) {
	t.Helper()

	value, exists := jsonObj[field]
	if !exists {
		t.Errorf("Field '%s' not found in JSON", field)
		return
	}

	if value != expected {
		t.Errorf("Incorrect field '%s': expected '%v', got '%v'", field, expected, value)
	}
}
