package handler

import (
	"fmt"
	"regexp"
	"strings"
)

func Solution(str string) []string {
	re := regexp.MustCompile("([a-zA-z]){2}")

	split := re.Split(str, -1)
	set := []string{}
	set = append(set, split...)
	return set
}

func ToCamelCase(s string) string {

	// dPost := []string{}
	// if strings.Contains(s, "-") {
	// 	dPost = strings.Split(s, "-")
	// } else {
	// 	dPost = strings.Split(s, "_")
	// }

	// result := make([]string, len(dPost))
	// for i, v := range dPost {
	// 	if len(v) > 0 {
	// 		if i != 0 {
	// 			if string(v[0]) != strings.ToUpper(string(v[0])) {
	// 				v = strings.ToUpper(string(v[0])) + v[1:]
	// 				result = append(result, v)
	// 			} else {
	// 				result = append(result, v)
	// 			}
	// 		} else {
	// 			result = append(result, v)
	// 		}
	// 	}
	// }

	// return strings.Join(result, "")

	return regexp.MustCompile("[-_](.)").ReplaceAllStringFunc(s, func(w string) string {
		return strings.ToUpper(w[1:])
	})
}

func FindOdd(seq []int) int {
	odds := make(map[int]int)

	for _, v := range seq {
		odds[v]++
	}

	for k, v := range odds {
		if v%2 != 0 {
			return k
		}
	}

	return 0
}
func WordCount(s string) map[string]int {
	cc := make(map[string]int)

	for _, v := range strings.Fields(s) {
		cc[v] = 1
	}

	return cc
}
func Fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func Fibonacci2(n int) int {
	first, second := 0, 1
	temp := 0
	for i := 0; i < n; i++ {
		temp = first
		first, second = second, first+second

	}
	return temp
}

type Vertex struct {
	X int
	Y int
}

func Point() {
	var p *int
	i := 42
	p = &i
	fmt.Println(p, *p)
}
