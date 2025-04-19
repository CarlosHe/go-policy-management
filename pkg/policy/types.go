package policy

import (
	"time"
)

type Action string

type Resource string

type Effect string

const (
	Allow Effect = "Allow"
	Deny  Effect = "Deny"
)

type ConditionOperator string

const (
	StringEquals              ConditionOperator = "StringEquals"
	StringNotEquals           ConditionOperator = "StringNotEquals"
	StringEqualsIgnoreCase    ConditionOperator = "StringEqualsIgnoreCase"
	StringNotEqualsIgnoreCase ConditionOperator = "StringNotEqualsIgnoreCase"
	StringLike                ConditionOperator = "StringLike"
	StringNotLike             ConditionOperator = "StringNotLike"

	NumericEquals            ConditionOperator = "NumericEquals"
	NumericNotEquals         ConditionOperator = "NumericNotEquals"
	NumericLessThan          ConditionOperator = "NumericLessThan"
	NumericLessThanEquals    ConditionOperator = "NumericLessThanEquals"
	NumericGreaterThan       ConditionOperator = "NumericGreaterThan"
	NumericGreaterThanEquals ConditionOperator = "NumericGreaterThanEquals"

	DateEquals            ConditionOperator = "DateEquals"
	DateNotEquals         ConditionOperator = "DateNotEquals"
	DateLessThan          ConditionOperator = "DateLessThan"
	DateLessThanEquals    ConditionOperator = "DateLessThanEquals"
	DateGreaterThan       ConditionOperator = "DateGreaterThan"
	DateGreaterThanEquals ConditionOperator = "DateGreaterThanEquals"

	Bool ConditionOperator = "Bool"
)

type ConditionKey string

type ConditionValue interface{}

type Condition struct {
	Operator ConditionOperator `json:"operator"`
	Key      ConditionKey      `json:"key"`
	Value    ConditionValue    `json:"value"`
}

type Statement struct {
	ID         string      `json:"id,omitempty"`
	Effect     Effect      `json:"effect"`
	Actions    []Action    `json:"actions"`
	Resources  []Resource  `json:"resources"`
	Conditions []Condition `json:"conditions,omitempty"`
}

type Policy struct {
	Version     string      `json:"version"`
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Statements  []Statement `json:"statements"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"`
}

const (
	PolicyVersion = "2023-01-01"
)
