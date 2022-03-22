package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func openfile() {
	f, err := os.Create("emp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

}

func repetition(st string) map[string]int {

	// using strings.Field Function
	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

func main() {
	//================================create the file============================
	openfile()
	// updatefile()

	//====================================read the file==========================
	b, err := ioutil.ReadFile("temp.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	//fmt.Println(b) // print the content as 'bytes'

	str := string(b) // convert content to a 'string'

	fmt.Println(str) // print the content as a 'string'

	//========================print index wise=================

	for index, element := range repetition(str) {
		fmt.Println(index, "=", element)
	}

}

