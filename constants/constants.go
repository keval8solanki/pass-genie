package constants

const (
	LOWER_CASE  = "l"
	UPPER_CASE  = "u"
	SPECIAL_CHAR = "s"
	NUMBER      = "n"
	LENGTH      = "L"
	FILE_NAME = "f"
)

var UPPER_CASE_ALPHABETS = [...]string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

var LOWER_CASE_ALPHABETS = [...]string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z",
}
var SPECIAL_CHARACTERS = [...]string{
	"~", "`", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "'", ";", ",", ".",
}

var BATCH_SIZE = 10000