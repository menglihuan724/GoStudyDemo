package main

import "fmt"

type matrix struct {
	x [3][3]int

}


func main(){
	myarry:=[3][3]int{
		{2,3,4},{5,6,7},{8,9,10},
	}
	mymatrix:=new(matrix)
	mymatrix.x=myarry
	a:=mymatrix.searchMatrix(5)
	fmt.Print(a)

}

func (param matrix)searchMatrix(target int) bool {
	row:=len(param.x)
	col:=len(param.x[0])
	start,end:=0,row*col-1
	for start!=end {
		mid:=(end-start)>>1
		mid_val:=param.x[mid/col][mid%col]
		if target>mid_val{
			fmt.Print("目标数比中间数大,初始数加一")
			start++
		}else if target<mid_val  {
			fmt.Print("目标数比中间数小,终点数减一")
			end--
		}else{
			fmt.Print("目标数比中间数相等")
			return true
		}
	}
	return false
}