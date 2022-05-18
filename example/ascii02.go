package main

import (
	// "net/http"
	"compress/gzip"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("http.ResponseWriter sample"))
}

func httph() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8089", nil)
}
func fileCreate() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	io.WriteString(writer, "io.MultiWriger example\n")
}

func zipCreate() {
	file, err := os.Create("test.txt.gz")
	if err != nil {
		panic(err)
	}
	writer := gzip.NewWriter(file)
	writer.Header.Name = "test.txt"
	writer.Write([]byte("gzip.Writer example\n"))
	writer.Close()
}

func main() {
	fileCreate()
	zipCreate()

	httph()
}
