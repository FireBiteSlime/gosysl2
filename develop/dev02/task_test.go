package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpacker(t *testing.T) {
	var testData = []struct {
		request string
		expect  string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
		{`qwe\4\5`, "qwe45"},
		{`qwe\45`, "qwe44444"},
		{`qwe\\5`, `qwe\\\\\`},
	}

	unpacker := &Char_Unpacker{}

	for i, piece := range testData {
		testname := fmt.Sprintf("Тест номер %d", i+1)
		t.Run(testname, func(t *testing.T) {
			res, err := unpacker.Unpack(piece.request)
			assert.NoError(t, err)
			assert.EqualValues(t, piece.expect, res)
		})
	}
}
