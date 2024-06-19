package util

var (
	EmptyLineByte          byte = byte('\n')
	WhiteSpaceByte         byte = byte(' ')
	ColonByte              byte = byte(':')
	CommaByte              byte = byte(',')
	LeftCurlyBraceByte     byte = byte('{')
	RightCurlyBraceByte    byte = byte('}')
	LeftSquareBracketByte  byte = byte('[')
	RightSquareBracketByte byte = byte(']')
	DoubleQuoteByte        byte = byte('"')
	BackSlashByte          byte = byte('\\')
	ForwardSlashByte       byte = byte('/')
)

var (
	LeftCurlyBrace   string = "{"
	RightCurlyBrace  string = "}"
	LeftSquareBrace  string = "["
	RightSquareBrace string = "]"
	Comma            string = ","
	Colon            string = ":"
	DoubleQuote      string = "\""
)

var BackSlashRune rune = rune('\\')

var (
	CheckSlice                  = []byte{LeftCurlyBraceByte, LeftSquareBracketByte, ColonByte, CommaByte, RightSquareBracketByte, RightCurlyBraceByte}
	AllowedCharsAfterEscapeChar = []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'}
)
