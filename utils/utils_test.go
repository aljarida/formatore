package utils

import (
	"testing"
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
		if IsValidIdentifer(str) == nil {
			t.Fatal()
		}
	}

	validIdentifiers := []string{"name1", "NAME1", "name_join"}
	for _, str := range validIdentifiers {
		if IsValidIdentifer(str) != nil {
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
		Integer, Integer,
		Real, Real,
		Text, Text,
		Real, Real,
		Real, Real,
		Text, Text,	
		Text, Text,
		Integer, Real, 
		Real, Real,
		Text, Text,
		Text, Text,
		Integer, Text,
		Integer, Real,
	}

	for i, value := range values {
		if res := InferType(value); res != expected[i] {
			t.Fatalf("Expected %s but inferred %s from '%s'.", expected[i], res, values[i])
		}
	}
}
