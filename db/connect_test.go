package db

import (
	"formatore/enums"
	"testing"
)

func TestConnectToDB(t *testing.T) {
	_, err := ConnectToDB(enums.TestDBName)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
