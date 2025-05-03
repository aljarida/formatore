package utils

import (
	"fmt"
	"formatore/enums"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Source: https://sqlite.org/datatype3.html
var validSQLite3Types = map[string]bool{
	enums.Blob: true,
	enums.Integer: true,
	enums.Real: true,
	enums.Text: true,
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

// Function takes arbitrary # of strings.
// Efficiently builds concatenated version.
func JoinStrings(strs ...string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()	
}

func JoinWithCommasSpaces(values []string) string {
	var builder strings.Builder
	for i, value := range values {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(value)
	}

	return builder.String()
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

func Assert(condition bool, msg string) {
	if !condition {
		panic(msg)
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
