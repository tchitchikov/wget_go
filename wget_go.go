package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var fileName = flag.String("filename", "output.txt", "supply the filename you wish to create")
	var requestURL = flag.String("url", "http://google.com", "supply the url you wish to download a file from")
	flag.Parse()

	fmt.Printf("Downloading %s\n", *requestURL)
	resp, err := http.Get(*requestURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(*fileName, res, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Complete!")
}
