package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Token struct {
	Header interface{}
	Claims interface{}
}

func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}
	token := MustParseToken(input)
	output, _ := json.MarshalIndent(token, "", "  ")
	fmt.Println(string(output))
}

func ParseToken(inputString string) (*Token, error) {
	parts := strings.Split(inputString, ".")
	if len(parts) != 3 {
		log.Fatal("token should have three parts")
	}

	token := &Token{}
	err := DecodePart(parts[0], &token.Header)
	if err != nil {
		return nil, err
	}

	err = DecodePart(parts[1], &token.Claims)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodePart(part string, result *interface{}) error {
	if l := len(part) % 4; l > 0 {
		part += strings.Repeat("=", 4-l)
	}

	decoded, err := base64.URLEncoding.DecodeString(part)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(decoded, &result); err != nil {
		return err
	}
	return nil
}

func MustParseToken(inputString string) *Token {
	token, err := ParseToken(inputString)
	if err != nil {
		log.Fatal(err)
	}
	return token
}
