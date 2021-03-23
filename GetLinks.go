package main

import (
    "regexp"
    "strings"
)

func GetLinks(strPtr *string) []string{
    re := regexp.MustCompile("href=\\\"\\/\\/i\\.4cdn\\.org\\/[a-z]+\\/[0-9]*\\.(jpg|png|jpeg)")

    strList := re.FindAllString((*strPtr),-1)


    var n int = 0

    for i,_ := range strList {
        if i%2==0 {
            strList[n] = "https:" + strings.TrimLeft(strList[i],"href=\"")
            n = n + 1
        }
    }
    strList = strList[0:n]

    return strList
}
