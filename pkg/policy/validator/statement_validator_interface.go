package validator

import "github.com/CarlosHe/go-policy-management/pkg/policy"

type IStatementValidator interface {
	ValidateStatement(statement policy.Statement, index int) []ValidationError
}
