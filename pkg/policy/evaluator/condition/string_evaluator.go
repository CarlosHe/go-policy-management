package condition

type StringEvaluator struct {
	patternMatcher PatternMatcher
}

func NewStringEvaluator(matcher PatternMatcher) *StringEvaluator {
	return &StringEvaluator{
		patternMatcher: matcher,
	}
}

func (e *StringEvaluator) Equals(contextValue, conditionValue interface{}) bool {
	cv, ok1 := contextValue.(string)
	cdv, ok2 := conditionValue.(string)
	if !ok1 || !ok2 {
		return false
	}
	return cv == cdv
}

func (e *StringEvaluator) NotEquals(contextValue, conditionValue interface{}) bool {
	return !e.Equals(contextValue, conditionValue)
}

func (e *StringEvaluator) Like(contextValue, conditionValue interface{}) bool {
	cv, ok1 := contextValue.(string)
	cdv, ok2 := conditionValue.(string)
	if !ok1 || !ok2 {
		return false
	}

	return e.patternMatcher.MatchesPattern(cv, cdv)
}

func (e *StringEvaluator) NotLike(contextValue, conditionValue interface{}) bool {
	return !e.Like(contextValue, conditionValue)
}
