package main

import "testing"

func TestTrimBody(t *testing.T) {
	_, err := TrimCurlyBraces("")
	if err == nil {
		t.Errorf("empty string with no curly braces should return an error")
	}
}

func TestTrimBodyWithOneCurlyBrace(t *testing.T){
	_, err := TrimCurlyBraces("}")
	if err == nil {
		t.Errorf("empty string with no front curly braces should return an error")
	}

	_, err = TrimCurlyBraces("{")
	if err == nil {
		t.Errorf("empty string with no end curly braces should return an error")
	}
}

func TestNormalString(t *testing.T) {
	s := "hello"
	err := ValidString(s,len(s) )
	if err != nil {
		t.Errorf("should not return error for normal string")
	}
}

func TestOneCtrlChar(t *testing.T) {
	s:= "hello\n"
	err := ValidString(s,len(s) )
	if err != nil {
		t.Errorf("should not return error for normal string with a newline")
	}
}

