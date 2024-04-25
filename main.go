package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/enchik0reo/calculator/internal/app"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func() {
		if err := writer.Flush(); err != nil {
			doPanic("unexpected error: can't write to stdout", err)
		}
	}()

	exp, err := getExpression(reader)
	if err != nil {
		doPanic("", err)
	}

	res, err := app.Calculate(exp)
	if err != nil {
		doPanic("", err)
	}

	writer.WriteString(res)
}

func getExpression(reader *bufio.Reader) ([]string, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("unexpected error: can't read from stdin, %v", err)
	}

	str = strings.Trim(str, "\n")
	str = strings.Trim(str, "\t")
	str = strings.Trim(str, "\r")

	vals := strings.Split(str, " ")

	return vals, nil
}

func doPanic(description string, err error) {
	if err != nil {
		if description != "" {
			panic(fmt.Sprintf("%s: %v", description, err))
		} else {
			panic(err)
		}
	}
}
