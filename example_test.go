package go_example

import (
	"testing"
)

func Test1(t *testing.T) {
	if 2 != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func Test2(t *testing.T)  {
	if 3 == 3 {
		t.Fail()
	}

	t.Error("Some other string {}" , 42)

	t.Log("foobar, log")
}