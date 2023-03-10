package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Oh I do like to be beside the seaside"
	fmt.Println(s)
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.Index(s, "the"))
	fmt.Println(strings.Replace(s, "seaside", "bar", -1))

}
