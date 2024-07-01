package util

// does lexical analysis over the input json to create lex tokens
func Lexer(content []byte) []string {
	var isInsideString bool = false
	var prevByte byte
	var currBytes []byte
	var lexResult []string
	for _, byt := range content {
		if byt != DoubleQuoteByte || prevByte == BackSlashByte {
			// all other cases
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
					// if current char is any of { [ : , ] }
					if len(currBytes) > 0 {
						// for cases like `true,` -> to add the `true` token since for (true | null | false) -> there is no end of token character
						lexResult = append(lexResult, string(currBytes))
						currBytes = nil
					}
					// add the current slice
					lexResult = append(lexResult, string(byt))
				} else {
					// for cases like (true | false | null)
					currBytes = append(currBytes, byt)
				}
			}
		} else {
			// double quote chars enclosing the string
			currBytes = append(currBytes, byt)
			if isInsideString {
        // if the double quote is at the end of string
				lexResult = append(lexResult, string(currBytes))
				currBytes = nil
			}
      // toggle end of string flag
			isInsideString = !isInsideString
		}
		prevByte = byt
	}
  // for cases if the file doesnot contain braces ex : null (./demo1.json)
  if len(currBytes) > 0 {
    lexResult = append(lexResult, string(currBytes))
  }
	return lexResult
}
