package main

import (
    "net/http"
    "os"
    "time"
    "fmt"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Useage: gochan [urlname]")
        os.Exit(1)
    }

    //Create a client with timeout of 20 seconds
    client := &http.Client{
        Timeout : 20 * time.Second,
    }

    str,err := GetSource(client, os.Args[1])

    if err != nil{
        panic(err)
    }

    imageLinks := GetLinks(str)
/*
    for i,_ := range imageLinks {
        fmt.Println("Link ",i+1," : ",imageLinks[i])
    }
*/
    fmt.Println("Starting to download images")
    err = DownloadImages(imageLinks,client)

    if err != nil{
        panic(err)
    }
}
