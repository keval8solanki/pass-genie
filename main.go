package main

import (
	"pass-genie/args"
	"pass-genie/helpers"
)

func main() {

	flags := args.FlagParser()
	operations := helpers.OperationBuilder(flags)
	password := helpers.PasswordBatch(operations, flags.PasswordLength)
	helpers.PasswordDisplay(password)
	helpers.PasswordSaver(flags, password)

}
