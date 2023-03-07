package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var url = "https://lospec.com/palette-list/load?colorNumberFilterType=any&colorNumber=8&tag=hardware&sortingType=default"

func main() {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)
	fmt.Println("Response:\n", responseString)
}
