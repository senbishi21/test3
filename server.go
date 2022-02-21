package main

import (
	"fmt"
	"net/http"
)


////////////////////////////
/////// HTTP Handler ///////
////////////////////////////

func testAPIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Encryption success\n")
}

func main() {
	http.HandleFunc("/test", testAPIHandler)
	http.ListenAndServe(":8080", nil)
}
