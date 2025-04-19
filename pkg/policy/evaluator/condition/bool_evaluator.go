package condition

type BoolEvaluator struct{}

func NewBoolEvaluator() *BoolEvaluator {
	return &BoolEvaluator{}
}

func (e *BoolEvaluator) Equals(contextValue, conditionValue interface{}) bool {
	cv, ok1 := contextValue.(bool)
	cdv, ok2 := conditionValue.(bool)
	if !ok1 || !ok2 {
		return false
	}
	return cv == cdv
}
