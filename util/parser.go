package util

func Parser(lexToken []string) ([]string, bool) {
  var result bool = false
  token := lexToken[0]
  if token == LeftCurlyBrace {
    lexToken = lexToken[1:]
    lexToken, result = IsValidObject(lexToken)
  } else if token == LeftSquareBrace {
    lexToken = lexToken[1:]
    lexToken, result = IsValidArray(lexToken)
  } else {
    result = IsValidString(token) || IsValidNumber(token) || IsValidBoolean(token) || IsValidNull(token)
    lexToken = lexToken[1:]
  }
	return lexToken, result
}
