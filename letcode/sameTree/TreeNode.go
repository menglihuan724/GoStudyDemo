
package main

import "fmt"

type Node struct {
	 value string
	 left *Node
	 right *Node
}

func main()  {
 	var  two Node
	var  one Node
	one.value="321"
	two.value="123"
	if (isSameTree(&one,&two)) {
		fmt.Println("相等")
	}else {
		fmt.Println("不相等")
	}

}
func isSameTree( a, b *Node)  bool{
	if(a==nil && b==nil){
		return true
	}
	if(a==nil||b==nil){
		return false
	}
	if(a.value == b.value){
		return isSameTree(a.left,b.left)&&isSameTree(a.right,b.right)
	}else {
		return false
	}

}