package main;

import (
	"fmt"
)

type My struct {
	a int;
}

func main() {
	// var obj My;
	// obj.a = 10;
	
	var age int
	fmt.Print("请输入年龄: ")
	fmt.Scan(&age);
	
	fmt.Printf("A: %d", age);

}