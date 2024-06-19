package util

import (
	"fmt"
	"strings"
)

// entry function for validator
func Validate(content []byte) bool {
	var result bool = false
	var lexToken []string = Lexer(content)
	lexToken = CleanUp(lexToken)
	_, result = Parser(lexToken)
	return result
}

// clean up the spaces around the braces
func CleanUp(lexToken []string) []string {
	lenLexToken := len(lexToken)
  var trimmedLexToken []string
	for i := 0; i < lenLexToken; i++ {
    tempStr := strings.TrimSpace(lexToken[i])
    if tempStr != "" {
      trimmedLexToken = append(trimmedLexToken, tempStr)
    }
	}
	return trimmedLexToken
}

func PrintResult(result bool) {
	if result {
		fmt.Printf("The file is a valid\n")
	} else {
		fmt.Printf("The file is invalid\n")
	}
}
