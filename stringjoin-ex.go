package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := []string{"this ", "is ","a ", "string ", "join ", "example ", "!"}
	fmt.Println(join1(s))
	fmt.Println(join2(s))
	fmt.Println(join3(s))

}

func join1(strs []string) string {
	var ret string
	for _, s := range strs {
		ret += s
	}
	return ret
}

func join2(strs []string) string {
	var buf bytes.Buffer
	for _, s := range strs {
		buf.WriteString(s)
	}

	return buf.String()
}

func join3(strs []string) string {
	var ret string
	ret = strings.Join(strs, "")
	return ret
}
