package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	model "forum.com/model"
)

var (
	ext = map[string]bool{
		"jpg": true,
		"png": true,
	}
)

func ParseImage(w http.ResponseWriter, r *http.Request) (string, error) {
	r.ParseMultipartForm(10 << 20)
	// 2. retrieve file
	file, handler, err := r.FormFile("photo")
	if err != nil {
		if err == http.ErrMissingFile {
			return "", nil
		}
		fmt.Println("Error Retrieving the File", err)
		return "", err
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	if !ValidFormat(handler.Filename) {
		return "", model.ErrImageFormat
	}

	// 3. write temporary file on our server
	tempFile, err := ioutil.TempFile("./assets/uploads", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	tempFile.Write(fileBytes)
	// 4. return result
	print(">> path: ", tempFile.Name())
	return "/" + tempFile.Name(), nil
}

func ValidFormat(fileName string) bool {
	if len(fileName) == 0 {
		return false
	}
	splitted := strings.Split(fileName, ".")
	last := splitted[len(splitted)-1]
	if _, ok := ext[last]; !ok {
		return false
	}
	return true

}

func DeleteImage(path string) error {
	fmt.Printf("sad")
	return nil
}
