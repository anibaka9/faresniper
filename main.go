package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	file, err := os.ReadFile("test.html")
	if err != nil {
		log.Fatal("err")
	}
	re := regexp.MustCompile(`data\(\) \{(.*?)mounted\(\)`)
	data := re.Find(file)
	fmt.Println(data)
}
