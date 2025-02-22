package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"io"
	"log"
)

func main() {
	body := []byte(`{"model":"mistral"}`)
	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(body))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	} 

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}
