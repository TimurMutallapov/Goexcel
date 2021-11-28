package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request ) {
	fmt.Fprint(w, "upload file")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("upload file: % + v\n", handler.Filename)
	fmt.Printf("File Size: x + v\n", handler.Size)
	fmt.Printf("upload File: % +v\n", handler.Header)

	tempFile, err := ioutil.TempFile("file", "upload-*.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes(){
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)

}
func main() {
	fmt.Printf("server start 8080\n")
	fmt.Printf("Go file Upload")
	setupRoutes()
}
