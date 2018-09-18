package main

import (
	"fmt"
	"strings"
)

func main(){
	var2:="terrymengli"
	fmt.Print(convert(var2,2))

}
func convert(var1 string,rows int)string{
	var2:=strings.Fields(var1)
	length:=len(var2)
	var3:=[rows]string{}
	var4:=0
	for var4<length{
		for var5:=0;var5<rows&&var4<length ;var5++  {
			var3[var5]=var3[var5]+var2[var4]
			var4++
		}
		for var5:=rows-2;var5>=1&&var4<length;var5-- {
			var3[var5]=var3[var5]+var2[var4]
			var4++
		}
	}
	for var5:=1;var5<len(var3);var5++  {
			var3[0]=var3[0]+var3[var5]
	}
	return var3[0]

}