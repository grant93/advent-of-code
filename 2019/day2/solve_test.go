package main

import (
	"bufio"
	"os"
	"testing"
)

func readFile(filename string) []string {
	var data []string
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func TestDayTwo(t *testing.T) {
	var tests = []struct {
		filename string
		expected [2]int
	}{
		{"input.txt", [2]int{6627023, 4019}},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			data := readFile(tt.filename)
			answer := solve(data)
			if answer[0] != tt.expected[0] || answer[1] != tt.expected[1] {
				t.Errorf("Got %+v, expected %+v", answer, tt.expected)
			}
		})
	}
}
