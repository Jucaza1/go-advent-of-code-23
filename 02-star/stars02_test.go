package main

import "testing"

func TestProcessLine(t *testing.T){
    test:="two3one2five"
    n:=ProcessLine([]byte(test))
    if n!=25{
        t.Fatal("expected 25 but got ",n)
    }

}
func TestReplaceParse(t *testing.T){
    test:= "onetwothreefourfivesixseveneightninezero"
    res:= string(ReplaceParse([]byte(test)))
    if res!= "1234567890"{
        t.Fatal("expected 1234567890 but got ",res)
    }
}
