package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var nums []int = make([]int, 0)
	var numsStr string
	fmt.Print("nums = ")
	fmt.Scanf("%s", &numsStr)
	var numsStrL = strings.Trim(numsStr, "[")
	var numsStrNoSquare = strings.Trim(numsStrL, "]")
	var items = strings.Split(numsStrNoSquare, ",")
	for i := 0; i < len(items); i++ {
		n, e := strconv.ParseInt(items[i], 10, 32)
		if e != nil {
			fmt.Println(e)
			continue
		}
		nums = append(nums, int(n))
	}

	var target int
	fmt.Print("target = ")
	fmt.Scanf("%d", &target)

	var results []int = make([]int, 0)
	for i := 0; i < len(nums); i++ {
		var minusVal = target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if minusVal == nums[j] {
				results = append(results, i)
				results = append(results, j)
			}
		}
	}

	fmt.Println(results)
}
