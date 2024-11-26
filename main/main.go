package main

import (
	"fmt"
	"gittt"
)
func main(){
	a:=99
	fmt.Printf("%d\n",a)
	b:=11111
	c:=888
	fmt.Printf("%d\n",gittt.Sum(gittt.Sum(a,b),c))
}