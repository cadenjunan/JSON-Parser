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

