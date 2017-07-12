package databaseUtility_test

import (
	"github.com/albertocubeddu/goUtility/database"
	"testing"
)

func TestRetrieveDSN(t *testing.T) {
	result, _ := databaseUtility.RetrieveDSN("user", "password", "db-hosting", "db-name")
	seedExpected := "user:password@tcp(db-hosting)/db-name"

	if result != seedExpected {
		t.Fail()
	}

	_, err := databaseUtility.RetrieveDSN("password", "db-hosting", "db-name")
	seedExpected = "user:password@tcp(db-hosting)/db-name"

	if err == nil {
		t.Fail()
	}

	result2, err2 := databaseUtility.RetrieveDSN("user", "", "db-hosting", "db-name")
	seedExpected = "user:@tcp(db-hosting)/db-name"

	if err2 != nil || result2 != seedExpected {
		t.Fail()
	}

}
