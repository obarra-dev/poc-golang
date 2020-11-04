package cbu

import (
	"io/ioutil"
	"path/filepath"
	"time"
)

var CBUUploadPath string = filepath.Join("uploads")

type CBUFileInfo struct {
	Name       string    `json:"name"`
	UploadDate time.Time `json:"uploadDate"`
}

func GetCBUFileInfo() ([]CBUFileInfo, error) {
	cbusFileInfo := make([]CBUFileInfo, 0)
	files, err := ioutil.ReadDir(CBUUploadPath)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		cbusFileInfo = append(cbusFileInfo, CBUFileInfo{Name: f.Name(), UploadDate: f.ModTime()})
	}
	return cbusFileInfo, nil
}
