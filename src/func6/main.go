package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	p := new(int)
	fmt.Println("*p = ", *p, ", p = ", p)
	*p = 29
	fmt.Println("*p = ", *p)
	var str string = "1213123"
	fmt.Println("*p = ", len(str))
	n, err := strconv.Atoi("12")
	fmt.Println("*p = ", n, err)

	str = strconv.Itoa(12345)
	fmt.Println("str = ", str)

	var bytes = []byte("hello go")
	fmt.Println("str = ", bytes)

	str1 := string([]byte{97, 98, 99})
	fmt.Println("str1 = ", str1)

	str2 := strconv.FormatInt(123, 2) // 2->8,16S
	fmt.Println("str2 = ", str2)

	vbool := strings.Contains("“seafood”", "foo") //true
	fmt.Println("vbool = ", vbool)
	vcout := strings.Count("“seafood”", "o") //true
	fmt.Println("vcout = ", vcout)
	fmt.Println(strings.EqualFold("“abc”", "“Abc”")) // true
	fmt.Println(strings.Index("NLT_abc", "abc"))     //4

	fmt.Println(strings.LastIndex("go golang", "go"))
}
