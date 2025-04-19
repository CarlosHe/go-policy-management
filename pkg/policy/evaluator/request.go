package evaluator

import (
	"github.com/CarlosHe/go-policy-management/pkg/policy"
)

type Request struct {
	Principal string
	Action    policy.Action
	Resource  policy.Resource
	Context   map[string]interface{}
}
