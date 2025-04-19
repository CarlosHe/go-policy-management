package condition

type PatternMatcher interface {
	MatchesPattern(input, pattern string) bool
}
