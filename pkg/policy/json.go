package policy

import (
	"encoding/json"
	"errors"
	"time"
)

func (p Policy) MarshalJSON() ([]byte, error) {
	type Alias Policy
	return json.Marshal(&struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at,omitempty"`
		Alias
	}{
		CreatedAt: p.CreatedAt.Format(time.RFC3339),
		UpdatedAt: func() string {
			if p.UpdatedAt.IsZero() {
				return ""
			}
			return p.UpdatedAt.Format(time.RFC3339)
		}(),
		Alias: Alias(p),
	})
}

func (p *Policy) UnmarshalJSON(data []byte) error {
	type Alias Policy
	aux := &struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.CreatedAt == "" {
		return errors.New("created_at is required")
	}
	createdAt, err := time.Parse(time.RFC3339, aux.CreatedAt)
	if err != nil {
		return errors.New("invalid created_at format, must be RFC3339")
	}
	p.CreatedAt = createdAt

	if aux.UpdatedAt != "" {
		updatedAt, err := time.Parse(time.RFC3339, aux.UpdatedAt)
		if err != nil {
			return errors.New("invalid updated_at format, must be RFC3339")
		}
		p.UpdatedAt = updatedAt
	}

	return nil
}

func FromJSON(jsonStr string) (Policy, error) {
	var p Policy
	err := json.Unmarshal([]byte(jsonStr), &p)
	return p, err
}

func FromJSONArray(jsonStr string) ([]Policy, error) {
	var p []Policy
	err := json.Unmarshal([]byte(jsonStr), &p)
	return p, err
}

func (p Policy) ToJSON() (string, error) {
	bytes, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (p Policy) ToJSONIndent() (string, error) {
	bytes, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type PolicyList struct {
	Policies []Policy `json:"policies"`
}

func FromJSONList(jsonStr string) ([]Policy, error) {
	var pl PolicyList
	err := json.Unmarshal([]byte(jsonStr), &pl)
	return pl.Policies, err
}

func ToJSONList(policies []Policy) (string, error) {
	pl := PolicyList{Policies: policies}
	bytes, err := json.Marshal(pl)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ToJSONListIndent(policies []Policy) (string, error) {
	pl := PolicyList{Policies: policies}
	bytes, err := json.MarshalIndent(pl, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
