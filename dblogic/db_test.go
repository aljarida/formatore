package dblogic

import (
	"testing"
)

func TestConnectToDB(t *testing.T) {
	_, err := ConnectToDB()
	if err != nil {
		t.Fatalf("%v", err)
	}
}
