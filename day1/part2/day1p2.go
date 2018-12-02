package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	//read values from file
	dataBufff, err := ioutil.ReadFile("../input.txt")
	data := string(dataBufff)

	//check for errors
	checkErr(err)

	//wtf is this syntax
	frequencies := map[int]bool{}

	//split string and add numbers until a repeat is found
	total := 0
	for {
		stringsSplit := strings.Split(data, "\r\n")
		for _, element := range stringsSplit {
			convertedStuff, err := strconv.ParseInt(element, 10, 64)
			checkErr(err)
			total += int(convertedStuff)
			// using a map keeps me from ruining the world with an array
			// and situations like this are the reason maps were made
			if frequencies[total] == true {
				fmt.Println(total)
				return
			} else {
				frequencies[total] = true
			}
		}
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
