package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	resp, _ := http.Get("https://www.google.com/")

	// Method 1
	// b := make([]byte, 99999)
	// n, _ := resp.Body.Read(b)
	// fmt.Println(string(b[:n]))

	// Method 2
	io.Copy(os.Stdout, resp.Body)

}
