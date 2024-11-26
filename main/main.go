package main

import (
	"fmt"
	"gittt"
)
func main(){
	a:=80
	fmt.Printf("%d\n",a)
	b:=27
	c:=41
	fmt.Printf("%d\n",gittt.Sum(gittt.Sum(c,b),a))
}