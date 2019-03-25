package example2

import (
	"testing"
)

func TestX(t *testing.T2) {

	if !true {
		t.Fatal("Fatal error");
	}

}

func TestY(t *testing.T)  {

	t.Log("Some logging")
	if false {
		t.FailNow()
	}

	t.Log(t.Name())
}