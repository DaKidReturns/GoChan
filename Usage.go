package main

import (
	"fmt"
	"os"
)

//Prints out the useage
//Usage()
func Usage() {
	fmt.Println("Usage: gochan [urlname]")
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Can't find the HOME directory, may behave in an unexpected way")
	}
	fmt.Println("Saves the images to", dir+"/Pictures/GoChan")
}

//Prints out a short intro about the application
//Intro()
func Intro() {
	fmt.Println("GoChan Ver 2.0")
}
