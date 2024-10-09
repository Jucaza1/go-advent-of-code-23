package main

import "testing"

func TestProcessLine(t *testing.T){
    test:="twone3one2fiveight"
    n:=processLine([]byte(test))
    if n!=28{
        t.Fatal("expected 28 but got ",n)
    }
}

