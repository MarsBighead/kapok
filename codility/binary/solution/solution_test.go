package solution

import (
	"fmt"
	"strconv"
	"testing"
)

var S string = "001111010101111"

func TestSolution(t *testing.T) {
	s := make([]byte, 4000)
	for i := 0; i <= len(s)-1; i++ {
		s[i] = s[i] ^ '1'
	}
	S = string(s)
	//x := regexp.MustCompile(`^[0]*`).ReplaceAllString(S, "")
	//fmt.Println("x=", x)
	fmt.Println(strconv.FormatInt(28, 2))
	S = strconv.FormatInt(28, 2)
	//fmt.Println("lower s=", string(s))
	fmt.Println("lower s=", S)
	num := Solution(S)
	fmt.Println("solution=", num)
}
func TestSolution1(t *testing.T) {
	s := make([]byte, 4000)
	for i := 0; i <= len(s)-1; i++ {
		s[i] = s[i] ^ '1'
	}
	S = string(s)
	S = "0" + strconv.FormatInt(28, 2)
	num := Solution1(S)
	fmt.Println("solution1=", num)
}
