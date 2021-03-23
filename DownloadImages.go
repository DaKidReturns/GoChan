package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func DownloadImages(folderName string, strLinks []string, client *http.Client) error {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	currentDir, err := os.Getwd()

	if err != nil {
		return err
	}

	saveImageDir := homeDir + "/Pictures/Gochan/" + folderName + "/"

	if _, err = os.Stat(saveImageDir); os.IsNotExist(err) {
		fmt.Println("Creating Directory")
		err := os.MkdirAll(saveImageDir, 0751)
		if err != nil {
			return err
		}
	}

	re := regexp.MustCompile("[0-9]+\\.(png|jpeg|jpg|svg)")

	/*err = os.Chdir(saveImageDir)

	  if err != nil {
	       return err
	  }*/

	for i, _ := range strLinks {
		//TODO: Make a function for this
		fileName := re.FindString(strLinks[i])

		if _, err := os.Stat(saveImageDir + fileName); !os.IsNotExist(err) {
			//fmt.Println(fileName,"Exists!!!")
			continue
		}

		f, err := os.Create(saveImageDir + fileName + ".tmp")

		if err != nil {
			return err
		}

		defer f.Close()

		request, err := http.NewRequest("GET", strLinks[i], nil)

		if err != nil {
			return err
		}

		fmt.Print("\nDownloading file: ", fileName)

		response, err := client.Do(request)

		if err != nil {
			return err
		}

		defer response.Body.Close()

		_, err = io.Copy(f, response.Body)

		if err != nil {
			return err
		}

		err = os.Rename(saveImageDir+fileName+".tmp", saveImageDir+fileName)

		if err != nil {
			return err
		}

	}

	err = os.Chdir(currentDir)

	if err != nil {
		return err
	}

	return nil

}
