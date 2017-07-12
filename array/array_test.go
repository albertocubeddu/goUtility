package arrayUtility_test

import "testing"
import (
	"github.com/albertocubeddu/goUtility/array"
)

func TestContain(t *testing.T) {
	if !arrayUtility.Contain("test", []string{"test"}) {
		t.Fail()
	}

	if !arrayUtility.Contain("123", []string{"hello","world","123"}) {
		t.Fail()
	}

	if arrayUtility.Contain("test", []string{"hello","world"}) {
		t.Fail()
	}

	if arrayUtility.Contain("123", []string{"hello","world"}) {
		t.Fail()
	}
}


func TestRemoveDuplicatesUnordered(t *testing.T) {
	seed :=[]string{"hello","world"}
	seedExpected := []string{"hello","world"}

	LoopTestRemoveDuplicatesUnordered(seed, seedExpected, t)

	seed = []string{"hello","hello","world","hello"}
	seedExpected = []string{"hello","world"}

	LoopTestRemoveDuplicatesUnordered(seed, seedExpected, t)

	seed = []string{"world","hello","world","hello"}
	seedExpected = []string{"hello","world"}

	LoopTestRemoveDuplicatesUnordered(seed, seedExpected, t)

	seed = []string{"test","123","hello","world","hello"}
	seedExpected = []string{"123","world","test","hello"}

	LoopTestRemoveDuplicatesUnordered(seed, seedExpected, t)



}


func LoopTestRemoveDuplicatesUnordered(seed []string, seedExpected []string, t *testing.T) {
	result := arrayUtility.RemoveDuplicatesUnordered(seed)

	if len(result) != len(seedExpected){
		t.Fail()
	}

	for _, v := range result {
		if !arrayUtility.Contain(v, seedExpected) {
			t.Fail()
		}
	}
}
