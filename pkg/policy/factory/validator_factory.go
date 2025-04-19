package factory

import (
	"github.com/CarlosHe/go-policy-management/pkg/policy/validator"
)

type IValidatorFactory interface {
	CreatePolicyValidator() validator.IPolicyValidator
	CreateFullValidator() validator.IFullValidatorInterface
}

type DefaultValidatorFactory struct{}

func NewValidatorFactory() *DefaultValidatorFactory {
	return &DefaultValidatorFactory{}
}

func (f *DefaultValidatorFactory) CreatePolicyValidator() validator.IPolicyValidator {
	return validator.NewDefaultValidator()
}

func (f *DefaultValidatorFactory) CreateFullValidator() validator.IFullValidatorInterface {
	return validator.NewDefaultValidator()
}
