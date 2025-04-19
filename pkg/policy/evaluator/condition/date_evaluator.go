package condition

import (
	"strconv"
	"time"
)

type DateEvaluator struct{}

func NewDateEvaluator() *DateEvaluator {
	return &DateEvaluator{}
}

func (e *DateEvaluator) Equals(contextValue, conditionValue interface{}) bool {
	cv, err1 := toTime(contextValue)
	cdv, err2 := toTime(conditionValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return cv.Equal(cdv)
}

func (e *DateEvaluator) NotEquals(contextValue, conditionValue interface{}) bool {
	return !e.Equals(contextValue, conditionValue)
}

func (e *DateEvaluator) GreaterThan(contextValue, conditionValue interface{}) bool {
	cv, err1 := toTime(contextValue)
	cdv, err2 := toTime(conditionValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return cv.After(cdv)
}

func (e *DateEvaluator) GreaterThanEquals(contextValue, conditionValue interface{}) bool {
	return e.GreaterThan(contextValue, conditionValue) || e.Equals(contextValue, conditionValue)
}

func (e *DateEvaluator) LessThan(contextValue, conditionValue interface{}) bool {
	cv, err1 := toTime(contextValue)
	cdv, err2 := toTime(conditionValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return cv.Before(cdv)
}

func (e *DateEvaluator) LessThanEquals(contextValue, conditionValue interface{}) bool {
	return e.LessThan(contextValue, conditionValue) || e.Equals(contextValue, conditionValue)
}

func toTime(value interface{}) (time.Time, error) {
	switch v := value.(type) {
	case time.Time:
		return v, nil
	case string:
		formats := []string{
			time.RFC3339,
			"2006-01-02",
			"2006-01-02T15:04:05",
			"2006-01-02 15:04:05",
		}

		for _, format := range formats {
			if t, err := time.Parse(format, v); err == nil {
				return t, nil
			}
		}
	}

	return time.Time{}, &strconv.NumError{Func: "toTime", Num: "", Err: strconv.ErrSyntax}
}
