package main

import (
	"fmt"
	"gittt"
)
func main(){
	a:=122
	fmt.Printf("%d\n",a)
	b:=58
	c:=41
	fmt.Printf("%d\n",gittt.Sum(gittt.Sum(c,b),a))
}
