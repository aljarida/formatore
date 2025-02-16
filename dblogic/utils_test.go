package dblogic

import (
	"testing"
)

func TestJoinStrings (t *testing.T) {
	a := "A"
	b := "B"
	c := "C"
	expected := a + b + c
	result := joinStrings(a, b, c)
	if expected != result {
		t.Fatal("Result does not match expected.")
	}

	expected = a
	result = joinStrings(a)
	if expected != result {
		t.Fatal("Result does not match expected.")
	}

	expected = ""
	result = joinStrings()
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
	result := joinWithCommasSpaces(values)
	if expected != result {
		t.Fatalf("Expected '%s' but got '%s'.", expected, result)
	}

	values = []string{}
	expected = ""
	result = joinWithCommasSpaces(values)
	if expected != result {
		t.Fatalf("Expected '%s' but got '%s'.", expected, result)
	}
}
