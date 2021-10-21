package main

import (
	"log"
	"os"
	"strconv"
)

var numberOfDays = 25

func main() {
	for i := 1; i <= numberOfDays; i++ {
		var dirName = "Day " + strconv.Itoa(i)
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create(dirName + "\\.gitkeep")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

	}
}