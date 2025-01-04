package helpers

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"pass-genie/args"
	"pass-genie/constants"
)

func RandomNumberGenerator(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func PasswordSaver(args args.Arg, password string) {
	if args.Filename == "" {
		return
	}

	file, err := os.Create(args.Filename)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, writeErr := file.WriteString(password)
	if writeErr != nil {
		fmt.Println("Failed to save the password to the file")
		return
	}
}

func OperationBuilder(args args.Arg) []string {
	operations := []string{}
	if args.IsLower {
		operations = append(operations, constants.LOWER_CASE)
	}

	if args.IsUpper {
		operations = append(operations, constants.UPPER_CASE)
	}

	if args.IsSpecial {
		operations = append(operations, constants.SPECIAL_CHAR)
	}

	if args.IsNumber {
		operations = append(operations, constants.NUMBER)
	}

	return operations
}

func calculateTotalBatchSize(passwordLength int) int {
	totalBatches := 1
	if passwordLength > constants.BATCH_SIZE {
		totalBatches = passwordLength / constants.BATCH_SIZE
		totalBatches = int(math.Floor(float64(totalBatches)))
	}
	return totalBatches

}


func PasswordBatch(operations []string, passwordLength int) string {


	totalBatches := calculateTotalBatchSize(passwordLength)

	if totalBatches == 1 {
		return PasswordBuilder(operations, passwordLength)
	}

	remainder := int(math.Remainder(float64(passwordLength), float64(constants.BATCH_SIZE)))

	if remainder > 0 {
		totalBatches += 1
	}

	var channels []chan string

	// Write to channels
	for i := 0; i < totalBatches; i++ {

		isLastBatch := i == totalBatches-1

		var passwordChunkLength int
		if isLastBatch && remainder > 0 {
			passwordChunkLength = remainder
		} else {
			passwordChunkLength = constants.BATCH_SIZE
		}

		channel := make(chan string)
		go func() {
			passwordChunk := PasswordBuilder(operations, passwordChunkLength)
			channel <- passwordChunk
		}()
		channels = append(channels, channel)
	}

	// Read from Channels
	var password string = readFromChannels(channels)
	return password
}

func readFromChannels(channels []chan string) string {
	// Read from Channels
	var str string
	var channelLength int = len(channels)

	for i := 0; i < channelLength; i++ {
		str += <-channels[i]
	}

	return str

}

func PasswordBuilder(operations []string, passwordLength int) string {
	operationLen := len(operations)
	upperCaseLen := len(constants.UPPER_CASE_ALPHABETS)
	lowerCaseLen := len(constants.LOWER_CASE_ALPHABETS)
	specialCharLen := len(constants.SPECIAL_CHARACTERS)

	var password string = ""
	for i := 0; i < passwordLength; i++ {
		operationIdx := RandomNumberGenerator(0, operationLen-1)
		operation := operations[operationIdx]

		if operation == constants.LOWER_CASE {
			letterIdx := RandomNumberGenerator(0, lowerCaseLen-1)
			letter := constants.LOWER_CASE_ALPHABETS[letterIdx]
			password = password + letter
			continue
		}

		if operation == constants.UPPER_CASE {
			letterIdx := RandomNumberGenerator(0, upperCaseLen-1)
			letter := constants.UPPER_CASE_ALPHABETS[letterIdx]
			password = password + letter
			continue
		}

		if operation == constants.SPECIAL_CHAR {
			letterIdx := RandomNumberGenerator(0, specialCharLen-1)
			letter := constants.SPECIAL_CHARACTERS[letterIdx]
			password = password + letter
			continue
		}

		if operation == constants.NUMBER {
			num := RandomNumberGenerator(0, 9)
			letter := fmt.Sprintf("%v", num)
			password = password + letter
			continue
		}
	}

	return password
}

func PasswordDisplay(password string) {
	fmt.Println("Password generated and copied to clipboard")
	fmt.Println("------------------------------------------")
	fmt.Println(password)
}




