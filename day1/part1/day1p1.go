package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	//read values from file
	data, err := ioutil.ReadFile("../input.txt")

	//check for errors because apparently we have to do that
	checkErr(err)

	//split string and add numbers
	result := addNums(string(data))

	//return total
	fmt.Print(result)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func addNums(in string) int {
	//Split
	stringsSplit := strings.Split(in, "\r\n")

	//parse each number into an int and add them
	total := 0
	for _, element := range stringsSplit {
		converted, err := strconv.ParseInt(element, 10, 64)
		checkErr(err)
		total += int(converted)
	}
	return total
}
