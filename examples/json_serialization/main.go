package main

import (
	"fmt"
	"time"

	"github.com/CarlosHe/go-policy-management/pkg/policy"
	"github.com/CarlosHe/go-policy-management/pkg/policy/evaluator"
	"github.com/CarlosHe/go-policy-management/pkg/policy/factory"
)

func main() {
	fmt.Println("=== Exemplo de Serialização JSON de Políticas ===")

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

	fmt.Println("\n=== Política para JSON ===")
	readPolicyJSON, err := readPolicy.ToJSONIndent()
	if err != nil {
		fmt.Printf("Erro ao converter política para JSON: %v\n", err)
		return
	}
	fmt.Println(readPolicyJSON)

	fmt.Println("\n=== JSON para Política ===")
	jsonStr := `{
		"version": "2023-01-01",
		"id": "p-custom-001",
		"name": "CustomPolicy",
		"description": "Política criada a partir de JSON",
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
		fmt.Printf("Erro ao converter JSON para política: %v\n", err)
		return
	}

	fmt.Printf("Política carregada do JSON:\n")
	fmt.Printf("  ID: %s\n", customPolicy.ID)
	fmt.Printf("  Nome: %s\n", customPolicy.Name)
	fmt.Printf("  Descrição: %s\n", customPolicy.Description)
	fmt.Printf("  Statements: %d\n", len(customPolicy.Statements))
	fmt.Printf("  Criada em: %s\n", customPolicy.CreatedAt.Format(time.RFC3339))

	fmt.Println("\n=== Lista de Políticas para JSON ===")
	policies := []policy.Policy{readPolicy, adminPolicy, customPolicy}
	policiesJSON, err := policy.ToJSONListIndent(policies)
	if err != nil {
		fmt.Printf("Erro ao converter lista de políticas para JSON: %v\n", err)
		return
	}
	fmt.Println(policiesJSON)

	fmt.Println("\n=== JSON para Lista de Políticas ===")
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
		fmt.Printf("Erro ao converter JSON para lista de políticas: %v\n", err)
		return
	}

	fmt.Printf("Políticas carregadas do JSON: %d\n", len(jsonPolicies))
	for i, p := range jsonPolicies {
		fmt.Printf("  Política %d: %s (%s)\n", i+1, p.Name, p.ID)
	}

	fmt.Println("\n=== Demonstração: Workflow de Carregamento de Políticas ===")
	fmt.Println("1. Carregar políticas de um arquivo JSON (simulado)")
	fmt.Println("2. Validar cada política antes de usar")
	fmt.Println("3. Criar um avaliador com as políticas carregadas")

	jsonContent := jsonListStr

	loadedPolicies, err := policy.FromJSONList(jsonContent)
	if err != nil {
		fmt.Printf("Erro ao carregar políticas: %v\n", err)
		return
	}

	validatorFactory := factory.NewValidatorFactory()
	validator := validatorFactory.CreatePolicyValidator()

	var validPolicies []policy.Policy
	for i, p := range loadedPolicies {
		errors := validator.Validate(p)
		if len(errors) > 0 {
			fmt.Printf("Política %d (%s) inválida:\n", i+1, p.ID)
			for _, err := range errors {
				fmt.Printf("  - %s: %s\n", err.Field, err.Message)
			}
			continue
		}
		validPolicies = append(validPolicies, p)
	}

	fmt.Printf("Políticas válidas: %d/%d\n", len(validPolicies), len(loadedPolicies))

	evaluatorFactory := factory.NewEvaluatorFactory()
	eval := evaluatorFactory.CreatePolicyEvaluator(validPolicies...)

	fmt.Println("Avaliador criado com sucesso com as políticas carregadas do JSON!")

	request := evaluator.Request{
		Principal: "user-123",
		Action:    "read",
		Resource:  "resource:json:document1",
		Context: map[string]interface{}{
			"user.role": "standard",
		},
	}

	result := eval.Evaluate(request)
	fmt.Printf("Avaliação de acesso: %v (%s)\n", result.Allowed, result.Reason)
}
