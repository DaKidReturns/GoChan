package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
    "github.com/DaKidReturns/gochan"
)

func main() {
	gochan.Intro()
	if len(os.Args) < 2 {
		gochan.Usage()
		os.Exit(1)
	}

	//Create a client with timeout of 20 seconds
	client := &http.Client{
		Timeout: 600 * time.Second,
	}

	str, err := gochan.GetSource(client, os.Args[1])

	if err != nil {
		panic(err)
	}

	imageLinks := gochan.GetImageLinks(str)

	folderName := gochan.GetFolderName(str)
	fmt.Println("FolderName: ", folderName)

	fmt.Println("Starting to download images")
	err = gochan.DownloadImages(folderName, imageLinks, client)

	if err != nil {
		panic(err)
	}

}
