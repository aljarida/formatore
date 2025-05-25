package utils

import (
	"fmt"
	"formatore/src/enums"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var validSQLite3Types = map[string]bool{
	enums.Integer: true,
	enums.Real:    true,
	enums.Text:    true,
}

// Determine if a given string is a reserved keyword.
func IsReserved(str string) bool {
	str = strings.ToUpper(str)
	_, ok := enums.ReservedKeywords[str]
	return ok
}

func IsNotReserved(str string) bool {
	return !IsReserved(str)
}

// Determine if given string is a valid SQLite3 type.
func IsValidType(str string) bool {
	_, ok := validSQLite3Types[str]
	return ok
}

// Determine if a string is a valid identifier.
func IsValidIdentifier(str string) error {
	if IsReserved(str) {
		return fmt.Errorf("Must not use reserved keywords; found '%s'.", str)
	}
	for i, r := range str {
		validNum := unicode.IsNumber(r) && i != 0
		validLetter := unicode.IsLetter(r) || r == '_'
		if !validNum && !validLetter {
			return fmt.Errorf("Must use alphanumeric characters or '_'; found '%c'.", r)
		}
	}
	return nil
}

func JoinStrArrWith(values []string, glue string) string {
	var builder strings.Builder
	for i, value := range values {
		if i > 0 {
			builder.WriteString(glue)
		}
		builder.WriteString(value)
	}

	return builder.String()
}

// Function takes arbitrary # of strings.
// Efficiently builds concatenated version.
func JoinStrings(strs ...string) string {
	return JoinStrArrWith(strs, "")
}

func JoinWithCommasSpaces(values []string) string {
	return JoinStrArrWith(values, ", ")
}

func Map[T any, R any](items []T, mapFn func(item T) R) []R {
	res := make([]R, len(items))
	for i, item := range items {
		res[i] = mapFn(item)
	}
	return res
}

var generateUnixTime = func() int64 {
	return time.Now().UTC().UnixNano()
}

func UnixTimestamp() string {
	timestamp := generateUnixTime()
	return strconv.FormatInt(timestamp, 10)
}

func IsPositiveInteger(s string) (int, bool) {
	n, err := strconv.Atoi(s)
	if err != nil || n <= 0 {
		return 0, false
	}
	return n, true
}

func Assert(condition bool, msg ...string) {
	if !condition {
		panic(JoinStrArrWith(msg, "...\n"))
	}
}

// Infer the type of input as either INTEGER, REAL, or TEXT. (BLOB, NULL excluded.)
func InferType(value string) string {
	if _, err := strconv.Atoi(value); err == nil {
		return enums.Integer
	}

	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return enums.Real
	}

	return enums.Text
}

func Has(arr []string, s string) bool {
	for _, t := range arr {
		if s == t {
			return true
		}
	}
	return false
}

func ParenthesizeFirstChar(s string) string {
	if s == "" {
		return s
	}

	var b strings.Builder
	b.Grow(len(s) + 2)
	b.WriteRune('(')
	b.WriteRune(rune(s[0]))
	b.WriteRune(')')
	if len(s) > 1 {
		b.WriteString(s[1:])
	}

	return b.String()
}

func PrettyColumnNameAsQuestion(s string) string {
	var b strings.Builder
	b.Grow(len(s) + len("[?]"))

	first := true
	for _, r := range s {
		switch r {
		case '_':
			b.WriteRune(' ')
		default:
			if first {
				b.WriteRune(unicode.ToUpper(r))
				first = false
			} else {
				b.WriteRune(r)
			}
		}
	}

	b.WriteString("[?]")
	return b.String()
}

func GetSortedKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}

func MaybeAppendDotCSV(s string) string {
	if strings.HasSuffix(s, ".csv") {
		return s
	}

	return s + ".csv"
}
