package main

import "testing"

func TestMain(t *testing.T) {
	if 100 != 100 {
		t.error("not equal")
	}
}
