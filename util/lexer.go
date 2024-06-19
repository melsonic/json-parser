package util

// does lexical analysis over the input json to create lex tokens
func Lexer(content []byte) []string {
	var isInsideString bool = false
	var prevByte byte
	var currBytes []byte
	var lexResult []string
	for _, byt := range content {
		if byt != DoubleQuoteByte || prevByte == BackSlashByte {
			if isInsideString {
				currBytes = append(currBytes, byt)
			} else {
				var present bool = false
				for _, cs := range CheckSlice {
					if cs == byt {
						present = true
						break
					}
				}
				if present {
					if len(currBytes) > 0 {
						lexResult = append(lexResult, string(currBytes))
						currBytes = nil
					}
					lexResult = append(lexResult, string(byt))
				} else {
					currBytes = append(currBytes, byt)
				}
			}
		} else {
			currBytes = append(currBytes, byt)
			if isInsideString {
				lexResult = append(lexResult, string(currBytes))
				currBytes = nil
			}
			isInsideString = !isInsideString
		}
		prevByte = byt
	}
	return lexResult
}
