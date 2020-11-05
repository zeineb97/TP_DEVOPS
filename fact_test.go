package main

import "testing"

func TestFact(t *testing.T) {

	r := fact(-1)
	if r != -1 {
		t.Error("Fact(-1) should be -1 ")
	}

	r = fact(0)
	if r != 1 {
		t.Error("Fact(0) should be 1 ")
	}

	r = fact(1)

	if r != 1 {
		t.Error("Fact(1) should be 1 ")
	}

	r = fact(5)

	if r != 120 {
		t.Error("Fact(0) should be 120 ")
	}
}
