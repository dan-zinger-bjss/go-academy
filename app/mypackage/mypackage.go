package stuff

import (
	// "errors"
	"fmt"
	"strconv"
	// "time"
)

var Name string = "Dan"

func IntArrToStrArr(intArr []int) []string{
	var strArr []string
	for index,value := range intArr{
		fmt.Println("index:", index)
		fmt.Println("value:", value)
		strArr = append(strArr, strconv.Itoa(value))
	}
	return strArr
}