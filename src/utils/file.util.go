package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func File2Base64(filePath string) (*string, error) {
	var base64String string
	if buf, err := ioutil.ReadFile(filePath); err != nil {
		fmt.Println("Error while encode file to base64: " + err.Error())
		return nil, err
	} else {
		switch http.DetectContentType(buf) {
		case "image/jpeg":
			base64String += "data:image/jpeg;base64,"
		case "image/png":
			base64String += "data:image/png;base64,"
		}
		base64String += base64.StdEncoding.EncodeToString(buf)
		return &base64String, nil
	}
}
