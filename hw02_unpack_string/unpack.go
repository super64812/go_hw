package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}

	if isNotValidString(str) {
		return "", ErrInvalidString
	}

	var result strings.Builder
	var prev rune

	for i, curr := range str {

		//indicates the last iteration
		lastIteration := len(str) == i+utf8.RuneLen(curr)

		//string unpacking
		switch {

		//repeat previous char {digit} times
		case unicode.IsDigit(curr) && unicode.IsLetter(prev):
			digit, _ := strconv.Atoi(string(curr))
			result.WriteString(strings.Repeat(string(prev), digit))

		//write two last chars on last iteration
		case unicode.IsLetter(curr) && unicode.IsLetter(prev) && lastIteration:
			result.WriteRune(prev)
			result.WriteRune(curr)

		//write last char on last iteration
		case unicode.IsLetter(curr) && lastIteration:
			result.WriteRune(curr)

		//write chars one by one
		case unicode.IsLetter(prev):
			result.WriteRune(prev)

		default:
			//no default options
		}
		prev = curr
	}

	return result.String(), nil
}

/*
validate input string using regexp, not valid if:
	1) string starts from digit
	2) string ends by digit
	3) number of iterations more than 9
*/
func isNotValidString(str string) bool {
	result, _ := regexp.MatchString("^[0-9]|[0-9]{2}|[0-9]+$", str)
	return result
}
