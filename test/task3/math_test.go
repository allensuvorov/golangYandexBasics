package math

import "testing"

func TestAddPositive(t *testing.T) {
	sum, err := Add(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}
	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestAddNegative(t *testing.T) {
	_, err := Add(-1, 2)
	if err == nil {
		t.Error("first arg negative - expected error not be nil")
	}
	_, err = Add(1, -2)
	if err == nil {
		t.Error("second arg negative - expected error not be nil")
	}
	_, err = Add(-1, -2)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}

func TestZeroArgs(t *testing.T) {
	_, err := Add(0, 2)
	if err == nil {
		t.Error("first arg zero - expected error not be nil")
	}

	_, err = Add(2, 0)
	if err == nil {
		t.Error("second arg zero - expected error not be nil")
	}

	_, err = Add(0, 0)
	if err == nil {
		t.Error("all arg zero - expected error not be nil")
	}

}
