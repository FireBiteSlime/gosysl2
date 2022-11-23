package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Unpacker interface {
	Unpack(string) (string, error)
}

type Char_Unpacker struct {
}

func (v *Char_Unpacker) Unpack(input string) (string, error) {

	if len(input) == 0 {
		return "", nil
	}

	if _, err := strconv.Atoi(string(input[0])); err == nil {
		return "", errors.New("некорректная строка")
	}

	str := strings.Split(input, "")
	result := ""

	for i := 0; i < len(str); i++ {

		if string(str[i]) == `\` {

			result += str[i+1]
			i++

		} else {

			if n, err := strconv.Atoi(string(str[i])); err == nil {

				for j := 0; j < n-1; j++ {
					result += string(str[i-1])
				}

			} else {
				result += str[i]
			}
		}

	}

	return result, nil
}

func main() {
	// strings := []string{"a4bc2d5e", "abcd", "45", `qwe\4\5`, `qwe\45`, `qwe\\5`}
	str := "45"
	unpacker := &Char_Unpacker{}
	fmt.Println(unpacker.Unpack(str))
}
