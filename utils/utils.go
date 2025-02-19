package utils

import (
	"fmt"
	"strings"
	"unicode"
	"time"
	"strconv"
)

// Supposedly complete list of SQLite keywords that need to be handled for identifier sanitization.
// Source 1: https://stackoverflow.com/questions/3373234/what-sqlite-column-name-can-be-cannot-be.
// Source 2: claude.ai.
var reservedKeywords = map[string]bool{
    "ABORT": true,
    "ACTION": true,
    "ADD": true,
    "AFTER": true,
    "ALL": true,
    "ALTER": true,
    "ALWAYS": true,
    "ANALYZE": true,
    "AND": true,
    "AS": true,
    "ASC": true,
    "ATTACH": true,
    "AUTOINCREMENT": true,
    "BEFORE": true,
    "BEGIN": true,
    "BETWEEN": true,
    "BY": true,
    "CASCADE": true,
    "CASE": true,
    "CAST": true,
    "CHECK": true,
    "COLLATE": true,
    "COLUMN": true,
    "COMMIT": true,
    "CONFLICT": true,
    "CONSTRAINT": true,
    "CREATE": true,
    "CROSS": true,
    "CURRENT": true,
    "CURRENT_DATE": true,
    "CURRENT_TIME": true,
    "CURRENT_TIMESTAMP": true,
    "DATABASE": true,
    "DEFAULT": true,
    "DEFERRABLE": true,
    "DEFERRED": true,
    "DELETE": true,
    "DESC": true,
    "DETACH": true,
    "DISTINCT": true,
    "DO": true,
    "DROP": true,
    "EACH": true,
    "ELSE": true,
    "END": true,
    "ESCAPE": true,
    "EXCEPT": true,
    "EXCLUDE": true,
    "EXCLUSIVE": true,
    "EXISTS": true,
    "EXPLAIN": true,
    "FAIL": true,
    "FILTER": true,
    "FIRST": true,
    "FOLLOWING": true,
    "FOR": true,
    "FOREIGN": true,
    "FROM": true,
    "FULL": true,
    "GENERATED": true,
    "GLOB": true,
    "GROUP": true,
    "GROUPS": true,
    "HAVING": true,
    "IF": true,
    "IGNORE": true,
    "IMMEDIATE": true,
    "IN": true,
    "INDEX": true,
    "INDEXED": true,
    "INITIALLY": true,
    "INNER": true,
    "INSERT": true,
    "INSTEAD": true,
    "INTERSECT": true,
    "INTO": true,
    "IS": true,
    "ISNULL": true,
    "JOIN": true,
    "KEY": true,
    "LAST": true,
    "LEFT": true,
    "LIKE": true,
    "LIMIT": true,
    "MATCH": true,
    "MATERIALIZED": true,
    "NATURAL": true,
    "NO": true,
    "NOT": true,
    "NOTHING": true,
    "NOTNULL": true,
    "NULL": true,
    "NULLS": true,
    "OF": true,
    "OFFSET": true,
    "ON": true,
    "OR": true,
    "ORDER": true,
    "OTHERS": true,
    "OUTER": true,
    "OVER": true,
    "PARTITION": true,
    "PLAN": true,
    "PRAGMA": true,
    "PRECEDING": true,
    "PRIMARY": true,
    "QUERY": true,
    "RAISE": true,
    "RANGE": true,
    "RECURSIVE": true,
    "REFERENCES": true,
    "REGEXP": true,
    "REINDEX": true,
    "RELEASE": true,
    "RENAME": true,
    "REPLACE": true,
    "RESTRICT": true,
    "RETURNING": true,
    "RIGHT": true,
    "ROLLBACK": true,
    "ROW": true,
    "ROWS": true,
    "SAVEPOINT": true,
    "SELECT": true,
    "SET": true,
    "TABLE": true,
    "TEMP": true,
    "TEMPORARY": true,
    "THEN": true,
    "TIES": true,
    "TO": true,
    "TRANSACTION": true,
    "TRIGGER": true,
    "UNBOUNDED": true,
    "UNION": true,
    "UNIQUE": true,
    "UPDATE": true,
    "USING": true,
    "VACUUM": true,
    "VALUES": true,
    "VIEW": true,
    "VIRTUAL": true,
    "WHEN": true,
    "WHERE": true,
    "WINDOW": true,
    "WITH": true,
    "WITHOUT": true,
}

// Source: https://sqlite.org/datatype3.html
var validSQLite3Types = map[string]bool{
	"BLOB": true,
	"INTEGER": true,
	"REAL": true,
	"TEXT": true,
}

// Determine if a given string is a reserved keyword.
func IsReserved(str string) bool {
	str = strings.ToUpper(str)
	_, ok := reservedKeywords[str]
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
func IsValidIdentifer(str string) error {
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

// TODO: Add unit test.
var generateUnixTime = func() int64 {
	return time.Now().UTC().UnixNano()
}
func UnixTimestamp() string {
	timestamp := generateUnixTime()
	return strconv.FormatInt(timestamp, 10)
}

// Infer the type of input as either INTEGER, REAL, or TEXT. (BLOB, NULL excluded.)
func InferType(value string) string {
	if _, err := strconv.Atoi(value); err == nil {
		return Integer
	}

	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return Real
	}

	return Text
}
