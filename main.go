package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	Intro()
	if len(os.Args) < 2 {
		Usage()
		os.Exit(1)
	}

	//Create a client with timeout of 20 seconds
	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	str, err := GetSource(client, os.Args[1])

	if err != nil {
		panic(err)
	}

	imageLinks := GetImageLinks(str)

	folderName := GetFolderName(str)
	fmt.Println("FolderName: ", folderName)

	fmt.Println("Starting to download images")
	err = DownloadImages(folderName, imageLinks, client)

	if err != nil {
		panic(err)
	}

}
