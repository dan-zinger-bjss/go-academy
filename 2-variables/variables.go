package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	cv7 := "3.14"
	if cv8, err := strconv.ParseFloat(cv7, 64); err == nil{
		pl(cv8)
	}
	cv9 := fmt.Sprintf("%f", 3.14)

	pl()


}