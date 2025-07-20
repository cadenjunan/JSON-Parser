package main

import (
	"fmt"
	"os"
	"strings"
)

func TrimCurlyBraces(payload string) (string, error){
	s, found := strings.CutPrefix(payload, "{")
	if !found {
		return "", fmt.Errorf("no curly brace at the front of the json body")
	}
	s, found = strings.CutSuffix(s, "}")
	if !found {
		return "", fmt.Errorf("no curly brace at the end of the json body")
	}
	return s, nil
}
func ValidControlChars(r rune) bool {
	return r == '"' || r == '/' || r == '\\' || r == 'r' || r == 't' || r == 'n' || r =='f'
}
func ValidString(value string, size int) error {
	for i, r := range value {
		if i > 0 && value[i-1] == '\\' {
			if !ValidControlChars(r)  {
				return fmt.Errorf("invalid control char after a reverse solidus")
			}
		}
	}
	return nil
}
func ValidValue(value string) error {
	size := len(value)
	// string
	if value[0] == '"' && value[size-1] == '"' {
		return ValidString(value, size)
	} else if value[0] == '{' && value[size-1] == '}' { // json body

	}else if  value[0] == '[' && value[size-1] == ']' {

	}
	
	return fmt.Errorf("invalid value: %s", value)
	// either number or boolean
}
func ValidKeyPair(keyPair string) error {
	keyPairSplit := strings.Split(keyPair, ":")
	size := len(keyPairSplit)
	if size == 1 || size > 2 {
		return fmt.Errorf("invalid key pair value, it should only contain one value per key")
	}
	key, value := strings.TrimSpace(keyPairSplit[0]), strings.TrimSpace(keyPairSplit[1])
	fmt.Printf("%s : %s\n", key, value)
	return nil
}
func ValidBody(jsonPayload string) error {
	size := len(jsonPayload)
	if jsonPayload[0] != '{' && jsonPayload[size-1] != '}' {
		return fmt.Errorf("no curly braces at the front or end of json body")
	}
	body, err := TrimCurlyBraces(jsonPayload)
	if err != nil {
		return err
	}
	keyPairs := strings.Split(body, ",")
	for _, keyPair := range keyPairs {
		ValidKeyPair(keyPair)
	}
	return nil
}

func main() {
	b,err := os.ReadFile("test.txt")
	if err != nil {
		panic(err.Error())
	}
	body := string(b)
	err = ValidBody(body)
	if err!= nil{
		panic(err.Error())
	}
}