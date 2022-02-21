package main

import (
	"fmt"
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)


////////////////////////////
/////// HTTP Handler ///////
////////////////////////////

func testAPIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Encryption success\n")
}

func main() {
	_init()
	http.HandleFunc("/test", testAPIHandler)
	http.ListenAndServe(":8080", nil)
}
