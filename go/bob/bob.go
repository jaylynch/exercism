/*
Package bob provides a simple implementation of a rude teenager.

Input is accepted in the form of:

	bob.Hey(string)
*/
package bob

import (
	"regexp"
	"strings"
)

var responses = map[string]string{
	"default": "Whatever.",
	"empty":   "Fine. Be that way!",
	"query":   "Sure.",
	"yelling": "Whoa, chill out!",
}

// Hey implements bob's basic response logic
func Hey(remark string) string {
	cleanRemark := strings.TrimSpace(remark)
	disposition := "default"
	hasLowercase, _ := regexp.Compile("[a-z]")
	hasUppercase, _ := regexp.Compile("[A-Z]")

	if len(cleanRemark) == 0 {
		disposition = "empty"
	} else if cleanRemark[len(cleanRemark)-1:] == "?" {
		disposition = "query"
	}

	if hasUppercase.MatchString(remark) && !hasLowercase.MatchString(remark) {
		disposition = "yelling"
	}

	return responses[disposition]
}
