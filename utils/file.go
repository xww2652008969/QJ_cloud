package utils

import (
	"os"
)

//
func Createfolder(dirname string) (error, bool) {
	err := os.MkdirAll(dirname, 0777)
	if err != nil {
		return nil, true
	} else {
		return err, false

	}
}

func Writefile(filename string, chadate []byte) (error, bool) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	defer f.Close()
	if err != nil {
		return err, false

	} else {
		_, err := f.Write(chadate)
		if err != nil {
			return err, false
		}
		if err != nil {
			return err, false
		}
		return nil, true
	}
}
func Readfile(filename string) (error, []byte) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	defer f.Close()
	if err != nil {
		return err, nil
	} else {
		var data []byte
		_, err := f.Read(data)
		if err != nil {
			return err, nil
		}
		if err != nil {
			return err, nil
		}
		return nil, data
	}
}
