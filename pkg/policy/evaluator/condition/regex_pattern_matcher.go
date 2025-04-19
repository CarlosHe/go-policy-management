package condition

import (
	"regexp"
	"strings"
)

type RegexPatternMatcher struct{}

func NewRegexPatternMatcher() *RegexPatternMatcher {
	return &RegexPatternMatcher{}
}

func (m *RegexPatternMatcher) MatchesPattern(input, pattern string) bool {
	regexPattern := strings.Replace(regexp.QuoteMeta(pattern), "\\*", ".*", -1)
	regex, err := regexp.Compile("^" + regexPattern + "$")
	if err != nil {
		return false
	}
	return regex.MatchString(input)
}
