package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	urlFlag := flag.String("url", "", "URL of the website to download")
	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("You must specify a URL to download")
		return
	}

	download(*urlFlag)
}

func download(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading:", err)
		return
	}
	defer response.Body.Close()

	file, err := os.Create("index.html")
	if err != nil {
		fmt.Println("Error while creating file:", err)
		return
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Error while saving file:", err)
		return
	}

	fmt.Println(url, "has been downloaded successfully")
}
