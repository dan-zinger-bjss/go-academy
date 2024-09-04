package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

func main(){
	var str1 string = "hello"
	str2 := "my"
	str3 := "name"
	str4 := "is"

	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		pl(str1, str2, str3, str4, name)
	} else {
		log.Fatal(err)
	}
}