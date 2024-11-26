package main

import (
	"fmt"
	"gittt"
)
func main(){
	a:=42
	fmt.Printf("%d\n",a)
	b:=11111
	c:=888
	d:=666
	fmt.Printf("%d %d\n",gittt.Sum(gittt.Sum(a,b),c),d)
}