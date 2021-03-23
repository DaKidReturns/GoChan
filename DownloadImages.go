package main

import(
    "io"
    "os"
    "net/http"
/*  "errors" */
    "regexp"
    "fmt"
)

func DownloadImages(strLinks []string,client *http.Client)  error{
    homeDir, err := os.UserHomeDir()

    if err != nil {
        return err
    }

    currentDir, err:= os.Getwd()

    if err != nil {
        return err
    }

    saveImageDir := homeDir+"/Pictures/Gochan"

    err = os.MkdirAll(saveImageDir,0751)

    if err != nil {
        return err
    }

    re := regexp.MustCompile("[0-9]+\\.(png|jpeg|jpg|svg)")


    err = os.Chdir(saveImageDir)

    if err != nil {
         return err
    }

    for i,_ := range strLinks {
//TODO: Make a funtion for this
        fileName := re.FindString(strLinks[i])
        f, err  := os.Create("./"+fileName+".tmp")

        if err != nil {
             return err
        }

        defer f.Close()

        request,err := http.NewRequest("GET",strLinks[i],nil)

        if err != nil {
             return err
        }

        fmt.Print("\nDownloading file: ", strLinks[i])

        response,err := client.Do(request)

        if err != nil {
             return err
        }

        defer response.Body.Close()

        _,err = io.Copy(f,response.Body)

        if err != nil {
            return err
        }

        err = os.Rename(fileName+".tmp",fileName)

        if err != nil {
            return err
        }

    }

    err = os.Chdir(currentDir)

    if err!=nil {
        return nil
    }

    return nil

}
