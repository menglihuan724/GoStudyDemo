package main

import ("fmt"
	"strings"
)

func main()  {
	b:="hellow word"
	c:=lastword(b)
	fmt.Print(c)
}
func lastword(org string)int{
	a := len(strings.TrimSpace(org))-strings.LastIndex(strings.TrimSpace(org)," ")-1
	return a
}