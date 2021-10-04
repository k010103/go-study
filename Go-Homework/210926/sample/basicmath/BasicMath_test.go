package basicmath

import (
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 1 + 3
	got := add(1, 3)
	if expected != got {
		t.Errorf("expected: %d got: %d", expected, got)
	}
}
func TestSub(t *testing.T) {
	expected := 1 - 3
	got := sub(1, 3)

	if expected != got {
		t.Errorf("expected: %d got: %d", expected, got)
	}
}
func TestMulti(t *testing.T) {
	expected := 1 * 3
	got := multi(1, 3)

	if expected != got {
		t.Errorf("expected: %d got: %d", expected, got)
	}
}
func TestMod(t *testing.T) {
	expected := 1 % 3
	got := mod(1, 3)

	if expected != got {
		t.Errorf("expected: %d got: %d", expected, got)
	}
}

func TestDiv(t *testing.T) {
	expected := 1 / 3
	got := div(1, 3)

	if expected != got {
		t.Errorf("expected: %d got: %d", expected, got)
	}
}
