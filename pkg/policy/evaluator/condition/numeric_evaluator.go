package condition

import (
	"strconv"
)

type NumericEvaluator struct{}

func NewNumericEvaluator() *NumericEvaluator {
	return &NumericEvaluator{}
}

func (e *NumericEvaluator) Equals(contextValue, conditionValue interface{}) bool {
	cv, err1 := toFloat64(contextValue)
	cdv, err2 := toFloat64(conditionValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return cv == cdv
}

func (e *NumericEvaluator) NotEquals(contextValue, conditionValue interface{}) bool {
	return !e.Equals(contextValue, conditionValue)
}

func (e *NumericEvaluator) LessThan(contextValue, conditionValue interface{}) bool {
	cv, err1 := toFloat64(contextValue)
	cdv, err2 := toFloat64(conditionValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return cv < cdv
}

func (e *NumericEvaluator) LessThanEquals(contextValue, conditionValue interface{}) bool {
	return e.LessThan(contextValue, conditionValue) || e.Equals(contextValue, conditionValue)
}

func (e *NumericEvaluator) GreaterThan(contextValue, conditionValue interface{}) bool {
	cv, err1 := toFloat64(contextValue)
	cdv, err2 := toFloat64(conditionValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return cv > cdv
}

func (e *NumericEvaluator) GreaterThanEquals(contextValue, conditionValue interface{}) bool {
	return e.GreaterThan(contextValue, conditionValue) || e.Equals(contextValue, conditionValue)
}

func toFloat64(value interface{}) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, &strconv.NumError{Func: "toFloat64", Num: "", Err: strconv.ErrSyntax}
	}
}
