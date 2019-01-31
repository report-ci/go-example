package example2

import (
	"testing"
)

func TestX(t *testing.T) {

	if true {
		t.Fatal("Fatal error");
	}

}

func TestY(t *testing.T)  {

	t.Log("Some logging")
	if !false {
		t.FailNow()
	}

	t.Log(t.Name())
}