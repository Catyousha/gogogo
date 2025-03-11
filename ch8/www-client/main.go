package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	URL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error in parsing:", err)
		return
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	request, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		fmt.Println("Get:", err)
		return
	}

	httpData, err := c.Do(request)
	if err != nil {
		fmt.Println("Error in Do():", err)
		return
	}
	defer httpData.Body.Close()

	fmt.Println("Status code:", httpData.Status)
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Println(string(header))

	contentType := httpData.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Charset:", characterSet[1])
	}

	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}

	// exercise 2: Modify wwwClient.go to save the HTML output to an external file
	// Create output file (using hostname as filename)
	hostname := URL.Hostname()
	outputFile := hostname + ".html"
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Copy the body to both the file and count the bytes
	length, err := io.Copy(file, httpData.Body)
	if err != nil {
		fmt.Println("Error saving content:", err)
		return
	}

	fmt.Printf("Content saved to %s\n", outputFile)
	fmt.Println("Total size:", length)
}