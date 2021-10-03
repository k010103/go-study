package palin

import (
	"testing"
)

func TestOne(t *testing.T) {
	expected := "yes"
	got := Palin("121")
	if expected != got {
		t.Errorf("expected: %s got: %s", expected, got)
	}
}

func TestTwo(t *testing.T) {
	expected := "no"
	got := Palin("1231")
	if expected != got {
		t.Errorf("expected: %s got: %s", expected, got)
	}
}
func TestThree(t *testing.T) {
	expected := "yes"
	got := Palin("12421")
	if expected != got {
		t.Errorf("expected: %s got: %s", expected, got)
	}
}
func TestFour(t *testing.T) {
	expected := ""
	got := Palin("0")
	if expected != got {
		t.Errorf("expected: %s got: %s", expected, got)
	}
}
