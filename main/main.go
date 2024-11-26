package main

import (
	"fmt"
	"gittt"
)
func main(){
	a:=20
	fmt.Printf("%d\n",a)
	b:=50
	c:=70
	fmt.Printf("%d\n",gittt.Sum(gittt.Sum(a,b),c))
}