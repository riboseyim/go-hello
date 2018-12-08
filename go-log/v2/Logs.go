package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("------hello world-----")

	var filename = "test1.log"
	logToFile(filename, "AAAAA")
	logToFile(filename, "BBBBB")
	logToFile(filename, "CCCCC")

}

func logToFile(filename string, str string) {

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	log.SetOutput(f)

	log.Println(str)
}
