package main

import (
	"regexp"
	"strings"
)

//GetImageLinks(*string) []string - The html is sarched for Image links
func GetImageLinks(strPtr *string) []string {
	//Input: pointer to the HTML code of the webpage
	//Output: A slice containing the links to be downloaded

	re := regexp.MustCompile("href=\\\"\\/\\/i\\.4cdn\\.org\\/[a-z]+\\/[0-9]*\\.(jpg|png|jpeg)")

	strList := re.FindAllString((*strPtr), -1)

	var n int = 0

	for i, _ := range strList {
		if i%2 == 0 {
			strList[n] = "https:" + strings.TrimLeft(strList[i], "href=\"")
			n = n + 1
		}
	}
	strList = strList[0:n]

	return strList
}

//GetFolderName(*string) string - This funtions scans the html to give the folder a meaningful name
func GetFolderName(strPtr *string) string {
	//Input: The pointer to the string containing the html code
	//Output: The folder name of the folder where the images are to be saved

	re := regexp.MustCompile("\\<span class=\"subject\"\\>[a-zA-Z0-9\\s]*</span\\>")
	forumName := re.FindString(*strPtr)
	/*fmt.Printf("Forum Name before printing: %c\n",rune((*strPtr)[forumName[1]-1]))*/
	forumName = strings.TrimRight(strings.TrimLeft(forumName, "<span class=\"subject\">"), "</span>")

	re = regexp.MustCompile("<div class=\"thread\" id=\"t[0-9]+\">")
	threadName := re.FindString(*strPtr)
	threadName = strings.TrimRight(strings.TrimLeft(threadName, "<div class=\"thread\" id=\"t"), "\">")

	/*if strings.ContainsAny(forumName,"abcdefghijklmnopqrstuvwxyz1234567890") {*/
	if forumName == "" {
		return threadName
	} else {
		return threadName + "-" + forumName
	}
}
