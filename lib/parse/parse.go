package parse

import (
	"strconv"
	"strings"
)

func ParseInt64FromInput(input []string) ([][]int64, error) {
	nums := [][]int64{}
	for _, s := range input {
		parts := strings.Fields(s)
		tmp := []int64{}
		for _, i := range parts {
			a, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, a)
		}
		nums = append(nums, tmp)
	}
	return nums, nil
}

func ParseUint64FromInput(input []string) ([][]uint64, error) {
	nums := [][]uint64{}
	for _, s := range input {
		parts := strings.Fields(s)
		tmp := []uint64{}
		for _, i := range parts {
			a, err := strconv.ParseUint(i, 10, 64)
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, a)
		}
		nums = append(nums, tmp)
	}
	return nums, nil
}
