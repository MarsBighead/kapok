package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import (
	"fmt"
	"regexp"
	"strconv"
)

// Solution soultion
// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func Solution(S string) int {
	length := len(S)
	var cnt int
	if length > 60 {
		bnums := []byte(S)
		length = len(bnums)
		var b0, b1 byte = 48, 49
		for i := 0; i < length; i++ {
			if bnums[i] == b0 {
				bnums = bnums[i+1:]
				length--
				i--
				fmt.Println("i=", i)
			} else {
				break
			}
		}
		fmt.Println(string(bnums))
		for length > 0 {
			if bnums[length-1] == b1 {
				bnums[length-1] = b0
			} else {
				bnums = bnums[:length-1]
				length--
			}
			if len(bnums) == 0 {
				break
			}
			cnt++

		}
	} else {
		num, _ := strconv.ParseInt(S, 2, 0)
		fmt.Println(S, num, "num=", num>>1, num)
		for num > 0 {
			var tmp int64
			if num%2 == 0 {
				tmp = num / 2
			} else {
				tmp = num - 1
			}
			num = tmp
			cnt++
		}

	}
	return cnt
}

//Solution1 imporved solution
func Solution1(S string) int {
	S = regexp.MustCompile(`^[0]+`).ReplaceAllString(S, "")
	//fmt.Println(S)
	var cnt int
	for _, chr := range S {
		if chr == '1' {
			cnt++
		}
	}
	return len(S) + cnt - 1
}
