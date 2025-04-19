package validator

import "github.com/CarlosHe/go-policy-management/pkg/policy"

type IPolicyFieldsValidator interface {
	ValidateFields(policy policy.Policy) []ValidationError
}
