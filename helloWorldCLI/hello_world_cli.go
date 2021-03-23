package main

import (
	"errors"
	"flag"
	"fmt"
)

// var nFlag = flag.String("m", "", "message for function")
var message string

func init() {
	flag.StringVar(&message, "m", "", "Type a message")
}

func main() {
	flag.Parse()

	result, err := setName(message)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func setName(hello string) (string, error) {
	if hello == "Hello World!" {
		return hello, nil
	} else {
		return "", errors.New("Your string should be equal to Hello World!")
	}
}
