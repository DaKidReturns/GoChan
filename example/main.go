package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
    "github.com/DaKidReturns/GoChan"
)

func main() {
	GoChan.Intro()
	if len(os.Args) < 2 {
		GoChan.Usage()
		os.Exit(1)
	}

	//Create a client with timeout of 20 seconds
	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	str, err := GoChan.GetSource(client, os.Args[1])

	if err != nil {
		panic(err)
	}

	imageLinks := GoChan.GetImageLinks(str)

	folderName := GoChan.GetFolderName(str)
	fmt.Println("FolderName: ", folderName)

	fmt.Println("Starting to download images")
	err = GoChan.DownloadImages(folderName, imageLinks, client)

	if err != nil {
		panic(err)
	}

}
