package validator

import "github.com/CarlosHe/go-policy-management/pkg/policy"

type IPolicyValidator interface {
	Validate(policy policy.Policy) []ValidationError
}
