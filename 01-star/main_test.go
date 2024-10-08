package main

import "testing"

func TestProcessLine(t *testing.T) {
	test := "abe30ad8j3iii1abxxd!!!"
	n := processLine([]byte(test))
	if n != 31 {
		t.Fatal("expected 31 but got ", n)
	}
}
