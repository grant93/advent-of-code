package helpers

import (
	"strconv"
	"strings"
)

func DeliminatedStringToIntSlice(data string, delim string) []int {
	var output []int
	tmp := strings.Split(data, delim)
	for _, a := range tmp {
		converted, _ := strconv.Atoi(a)
		output = append(output, converted)
	}
	return output 
}
