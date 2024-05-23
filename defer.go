package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Defer file closure
	defer func() {
		log.Println("Closing the file.")
		file.Close()
	}()

	// Stacked defer. LIFO
	for i := 0; i < 5; i++ {
		defer log.Println(i)
	}

	log.Println("Writing to file.")
	_, err = file.WriteString("Golang.\n")
	if err != nil {
		log.Fatal(err)
	}
}
