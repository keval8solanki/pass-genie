package args

import (
	"flag"
	"pass-genie/constants"
)

type Arg struct {
	IsLower        bool
	IsUpper        bool
	IsSpecial      bool
	IsNumber       bool
	PasswordLength int
	Filename       string
}

func FlagParser() Arg {
	isLower := flag.Bool(constants.LOWER_CASE, true, "Include lower case characters")
	isUpper := flag.Bool(constants.UPPER_CASE, true, "Include upper case characters")
	isSpecial := flag.Bool(constants.SPECIAL_CHAR, true, "Include special characters")
	isNumber := flag.Bool(constants.NUMBER, true, "Include numbers")
	filename := flag.String(constants.FILE_NAME, "", "Password will be saved in this .txt file")
	passwordLength := flag.Int(constants.LENGTH, 16, "Length of the password")

	flag.Parse()

	return Arg{
		IsLower:        *isLower,
		IsUpper:        *isUpper,
		IsSpecial:      *isSpecial,
		IsNumber:       *isNumber,
		PasswordLength: *passwordLength,
		Filename:       *filename,
	}
}





