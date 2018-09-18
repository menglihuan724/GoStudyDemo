package main

import "fmt"

func main()  {
	var a=500
	fmt.Print(reverse(a))
}

func reverse(x int) int {
	var result int= 0;

	for 
	{
		var  tail int= x % 10
		var newResult int= result * 10 + tail
		/*检测溢出*/
		var test = (newResult - tail) / 10
		if test!= result{
			return 0
		}
		result = newResult
		x = x / 10
		if x==0 {
			break
		}
	}

	return result;
}
