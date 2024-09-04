package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"reflect"
)

func main(){
	fmt.Println("Hello", stuff.Name)
	var iArr = [5]int{3,4,5,6,7}
	thing := stuff.IntArrToStrArr(iArr[:])
	fmt.Println((thing))

	fmt.Println(reflect.TypeOf(thing))
}