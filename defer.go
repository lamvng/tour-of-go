package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Opening the file...")
	// file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE, 0644)
	file, err := os.OpenFile("example.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Defer the closure of the file
	defer func() {
		fmt.Println("Closing the file")
		file.Close()
	}()

	fmt.Println("Performing some operations on the file")
	// Perform some operations on the file

	fmt.Println("Function is about to end")
}
