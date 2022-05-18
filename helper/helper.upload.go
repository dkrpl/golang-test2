package helpers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"golang-test2/config"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
)

func Upload(collection, name, img string) (string, error) {
	path := ""
	if img == "" {
		path = "/picture/" + collection + "/" + name + ".png"
		return path, nil
	}
	path = config.GetDir() + "/assets/picture/" + collection
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
	idx := strings.Index(img, ";base64,")
	if idx < 0 {
		err := errors.New("error indexing image")
		return "", err
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(img[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(path+"/"+name+".png", buff.Bytes(), 0644)
	if err != nil {
		return "", err
	}
	path = "/picture/" + collection + "/" + name + ".png"

	return path, nil
}

func UploadWithReplace(path, image string) error {
	path = config.GetDir() + "/assets" + path
	idx := strings.Index(image, ";base64,")
	if idx < 0 {
		err := errors.New("error indexing image")
		return err
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(image[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, buff.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}

func UploadWithReplaceV2(collection string, data *multipart.FileHeader) (path string, err error) {
	file, err := data.Open()
	defer file.Close()

	if err != nil {
		return "", err
	}
	path = config.GetDir() + "/assets/picture/" + collection
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
	// fileData, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	return "", err
	// }
	outFile, err := os.Create(path + "/" + data.Filename)
	if err != nil {
		return "", err
	}
	io.Copy(outFile, file)
	path = "/picture/" + collection + "/" + data.Filename
	return path, nil
}
