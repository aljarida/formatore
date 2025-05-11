package utils

import (
	"time"
	"testing"
	"formatore/enums"
	"github.com/stretchr/testify/assert"
)

func TestJoinStrings (t *testing.T) {
	a := "A"
	b := "B"
	c := "C"
	expected := a + b + c
	result := JoinStrings(a, b, c)
	if expected != result {
		t.Fatal("Result does not match expected.")
	}

	expected = a
	result = JoinStrings(a)
	if expected != result {
		t.Fatal("Result does not match expected.")
	}

	expected = ""
	result = JoinStrings()
	if expected != result {
		t.Fatal("Result does not match expected.")
	}
}

func TestIsReserved(t *testing.T) {
	expected := true
	result := IsReserved("TABLE")
	if expected != result {
		t.Fatal()
	}
	
	expected = false
	result = IsReserved("DOLLARS")
	if expected != result {
		t.Fatal()
	}
}

func TestIsValidType(t *testing.T) {
	expected := true
	result := IsValidType("TEXT")
	if expected != result {
		t.Fatal()
	}

	expected = false
	result = IsValidType("DATETIME")
	if expected != result {
		t.Fatal()
	}
}

func TestIsValidIdentifer(t *testing.T) {
	invalidIdentifers := []string{"1name", "join", "join$"}
	for _, str := range invalidIdentifers {
		if IsValidIdentifier(str) == nil {
			t.Fatal()
		}
	}

	validIdentifiers := []string{"name1", "NAME1", "name_join"}
	for _, str := range validIdentifiers {
		if IsValidIdentifier(str) != nil {
			t.Fatal()
		}
	}
}

func TestJoinWithCommasSpaces(t *testing.T) {
	values := []string{"1", "2", "3"}
	expected := "1, 2, 3"
	result := JoinWithCommasSpaces(values)
	if expected != result {
		t.Fatalf("Expected '%s' but got '%s'.", expected, result)
	}

	values = []string{}
	expected = ""
	result = JoinWithCommasSpaces(values)
	if expected != result {
		t.Fatalf("Expected '%s' but got '%s'.", expected, result)
	}
}

func TestInferType(t *testing.T) {
	values := []string{
		"1", "-1",
		"1.", "-1.",
		" 1 ", " - 1 ",
		"1.1", "-1.1",
		".1", "-.1",
		" . 1 ", " - . 1 ",
		"--1", "+ 1.1",
		"+1", "+1.1",
		"11111.1", "1.11111",
		"1 . 1", "1 . 1.",
		".", "$",
		"0", "0,1",
		"00", "0.0",
	}

	expected := []string{
		enums.Integer, enums.Integer,
		enums.Real, enums.Real,
		enums.Text, enums.Text,
		enums.Real, enums.Real,
		enums.Real, enums.Real,
		enums.Text, enums.Text,	
		enums.Text, enums.Text,
		enums.Integer, enums.Real, 
		enums.Real, enums.Real,
		enums.Text, enums.Text,
		enums.Text, enums.Text,
		enums.Integer, enums.Text,
		enums.Integer, enums.Real,
	}

	for i, value := range values {
		if res := InferType(value); res != expected[i] {
			t.Fatalf("Expected %s but inferred %s from '%s'.", expected[i], res, values[i])
		}
	}
}

func TestMap(t *testing.T) {
	addOne := func(i int) int { return i + 1 }

	arr1 := []int{1,2,3,4,5}
	expected1 := []int{2,3,4,5,6}
	actual1 := Map(arr1, addOne)
	assert.Equal(t, expected1, actual1, "Should be equal.")

	arr2 := []int{}	
	expected2 := []int{}
	actual2 := Map(arr2, addOne)
	assert.Equal(t, expected2, actual2, "Should be equal.")
}

func TestTime(t *testing.T) {
	start := time.Now().UTC().UnixNano()
	got := generateUnixTime()
	end := time.Now().UTC().UnixNano()

	assert.Greater(t, got, start, "Expected acquired time to be greater than start time.")
	assert.Greater(t, end, got, "Expected end time to be greater than acquired time.")
}

func TestParenthesizeFirstChar(t *testing.T) {
	s := ParenthesizeFirstChar("T")
	assert.Equal(t, "(T)", s)

	s = ParenthesizeFirstChar("")
	assert.Equal(t, "", s)
}

func TestPrettyColumnNameAsQuestion(t *testing.T) {
	s := PrettyColumnNameAsQuestion("a_b_c_d")
	assert.Equal(t, "A b c d[?]", s)
	
	s = PrettyColumnNameAsQuestion("")
	assert.Equal(t, "[?]", s)
}
