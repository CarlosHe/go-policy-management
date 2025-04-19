package factory

import (
	"time"

	"github.com/CarlosHe/go-policy-management/pkg/policy"
)

type IPolicyFactory interface {
	CreatePolicy(id, name string, statements ...policy.Statement) policy.Policy
	CreateStatement(id string, effect policy.Effect, actions []policy.Action, resources []policy.Resource) policy.Statement
}

type DefaultPolicyFactory struct{}

func NewPolicyFactory() *DefaultPolicyFactory {
	return &DefaultPolicyFactory{}
}

func (f *DefaultPolicyFactory) CreatePolicy(id, name string, statements ...policy.Statement) policy.Policy {
	return policy.Policy{
		Version:    policy.PolicyVersion,
		ID:         id,
		Name:       name,
		Statements: statements,
		CreatedAt:  time.Now(),
	}
}

func (f *DefaultPolicyFactory) CreateStatement(id string, effect policy.Effect, actions []policy.Action, resources []policy.Resource) policy.Statement {
	return policy.Statement{
		ID:        id,
		Effect:    effect,
		Actions:   actions,
		Resources: resources,
	}
}
