package evaluator

import (
	"time"
)

type Result struct {
	Allowed      bool
	Reason       string
	EvaluatedAt  time.Time
	MatchedRules []string
}
