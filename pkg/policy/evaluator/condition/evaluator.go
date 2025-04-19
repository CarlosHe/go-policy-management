package condition

import (
	"github.com/CarlosHe/go-policy-management/pkg/policy"
)

type Evaluator interface {
	Evaluate(condition policy.Condition, context map[string]interface{}) bool
}
