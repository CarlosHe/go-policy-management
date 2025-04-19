# Go Policy Management

[![Go Report Card](https://goreportcard.com/badge/github.com/CarlosHe/go-policy-management)](https://goreportcard.com/report/github.com/CarlosHe/go-policy-management)
[![GoDoc](https://godoc.org/github.com/CarlosHe/go-policy-management?status.svg)](https://godoc.org/github.com/CarlosHe/go-policy-management)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A powerful and flexible policy management library for Go applications. This library provides a robust framework for defining, validating, and evaluating access control policies in your Go applications.

## Features

- **Attribute-Based Access Control (ABAC)** - Define fine-grained policies based on attributes
- **Policy Validation** - Validate policy structure before enforcement
- **Flexible Conditions** - Support for various condition types (string, numeric, date, boolean)
- **Extensible Design** - Easy to extend with custom validators and evaluators
- **SOLID Architecture** - Built with SOLID principles for maintainability and extensibility
- **JSON Serialization** - Convert policies to/from JSON for storage and transmission

## Installation

```bash
go get github.com/CarlosHe/go-policy-management
```

## Quick Start

### 1. Define a Policy

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/CarlosHe/go-policy-management/pkg/policy"
    "github.com/CarlosHe/go-policy-management/pkg/policy/factory"
)

func main() {
    // Create factories
    policyFactory := factory.NewPolicyFactory()
    
    // Create a read-only policy
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
    
    fmt.Printf("Created policy: %s\n", readOnlyPolicy.Name)
}
```

### 2. Validate a Policy

```go
package main

import (
    "fmt"
    
    "github.com/CarlosHe/go-policy-management/pkg/policy"
    "github.com/CarlosHe/go-policy-management/pkg/policy/factory"
)

func main() {
    // Create factories
    policyFactory := factory.NewPolicyFactory()
    validatorFactory := factory.NewValidatorFactory()
    
    // Create a policy
    myPolicy := policyFactory.CreatePolicy(
        "p-1234",
        "MyPolicy",
        policyFactory.CreateStatement(
            "s-1",
            policy.Allow,
            []policy.Action{"read"},
            []policy.Resource{"resource:user:*"},
        ),
    )
    
    // Validate the policy
    validator := validatorFactory.CreatePolicyValidator()
    errors := validator.Validate(myPolicy)
    
    if len(errors) > 0 {
        fmt.Println("Validation errors:")
        for _, err := range errors {
            fmt.Printf("- %s: %s\n", err.Field, err.Message)
        }
        return
    }
    
    fmt.Println("Policy is valid!")
}
```

### 3. Evaluate Policies for Authorization

```go
package main

import (
    "fmt"
    
    "github.com/CarlosHe/go-policy-management/pkg/policy"
    "github.com/CarlosHe/go-policy-management/pkg/policy/evaluator"
    "github.com/CarlosHe/go-policy-management/pkg/policy/factory"
)

func main() {
    // Create factories
    policyFactory := factory.NewPolicyFactory()
    evaluatorFactory := factory.NewEvaluatorFactory()
    
    // Create a policy
    readOnlyPolicy := policyFactory.CreatePolicy(
        "p-1234",
        "ReadOnlyAccess",
        policyFactory.CreateStatement(
            "s-1",
            policy.Allow,
            []policy.Action{"read"},
            []policy.Resource{"resource:*"},
        ),
    )
    
    // Create a policy evaluator
    eval := evaluatorFactory.CreatePolicyEvaluator(readOnlyPolicy)
    
    // Evaluate a request
    request := evaluator.Request{
        Principal: "user-123",
        Action:    "read",
        Resource:  "resource:doc1",
        Context: map[string]interface{}{
            "time": "2023-05-01T10:00:00Z",
        },
    }
    
    result := eval.Evaluate(request)
    
    fmt.Printf("Authorization result: %v\n", result.Allowed)
    fmt.Printf("Reason: %s\n", result.Reason)
}
```

### 4. JSON Serialization

```go
package main

import (
    "fmt"
    
    "github.com/CarlosHe/go-policy-management/pkg/policy"
    "github.com/CarlosHe/go-policy-management/pkg/policy/factory"
)

func main() {
    // Create a policy
    policyFactory := factory.NewPolicyFactory()
    myPolicy := policyFactory.CreatePolicy(
        "p-1234",
        "MyPolicy",
        policyFactory.CreateStatement(
            "s-1",
            policy.Allow,
            []policy.Action{"read"},
            []policy.Resource{"resource:*"},
        ),
    )
    
    // Convert policy to JSON
    jsonStr, err := myPolicy.ToJSON()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Policy as JSON: %s\n", jsonStr)
    
    // Convert JSON to policy
    jsonPolicy := `{
        "version": "2023-01-01",
        "id": "p-json-001",
        "name": "JSONPolicy",
        "statements": [
            {
                "id": "s-1",
                "effect": "Allow",
                "actions": ["read"],
                "resources": ["resource:document:*"]
            }
        ],
        "created_at": "2023-05-01T10:00:00Z"
    }`
    
    loadedPolicy, err := policy.FromJSON(jsonPolicy)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Loaded policy: %s (ID: %s)\n", loadedPolicy.Name, loadedPolicy.ID)
}
```

## Advanced Usage

### Policies with Conditions

```go
// Creating a policy with conditions
adminPolicy := policyFactory.CreatePolicy(
    "p-admin",
    "AdminAccess",
    policyFactory.CreateStatement(
        "s-1",
        policy.Allow,
        []policy.Action{"*"},
        []policy.Resource{"*"},
    ),
)

// Add conditions to the statement
adminPolicy.Statements[0].Conditions = []policy.Condition{
    {
        Operator: policy.StringEquals,
        Key:      "user.role",
        Value:    "admin",
    },
}
```

### Custom Factories

You can create custom factories by implementing the interfaces:

```go
// Example of a custom condition factory
type MyConditionFactory struct {
    factory.DefaultConditionFactory
}

func (f *MyConditionFactory) CreatePatternMatcher() condition.PatternMatcher {
    return &MyCustomPatternMatcher{}
}
```

### Loading Policies from Storage

```go
// Load policies from JSON file
jsonData, err := ioutil.ReadFile("policies.json")
if err != nil {
    log.Fatalf("Failed to read policies file: %v", err)
}

// Parse JSON into policies
policies, err := policy.FromJSONList(string(jsonData))
if err != nil {
    log.Fatalf("Failed to parse policies: %v", err)
}

// Create evaluator with loaded policies
evaluator := evaluatorFactory.CreatePolicyEvaluator(policies...)
```

## Architecture

The library is designed with SOLID principles:

- **Single Responsibility Principle**: Each class has only one responsibility
- **Open/Closed Principle**: Components are open for extension but closed for modification
- **Liskov Substitution Principle**: Implementations can be substituted without affecting correctness
- **Interface Segregation Principle**: Small, focused interfaces instead of large, monolithic ones
- **Dependency Inversion Principle**: High-level modules depend on abstractions, not on details

### Core Components

- **Policy**: The main entity representing access control rules
- **Statement**: Individual rule within a policy
- **Condition**: Additional constraints that can be applied to statements
- **Validator**: Ensures policies are well-formed
- **Evaluator**: Determines if a request matches policy rules

## Examples

For more examples, check the `examples/` directory in this repository, including:

- Basic policy usage
- Advanced policy configurations
- JSON serialization and deserialization
- Custom condition implementations
- Role-based access control

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 