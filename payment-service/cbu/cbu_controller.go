package cbu

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"poc-golang/payment-service/webservices/middleware"
)

const cbusPath = "cbus"

// SetupRoutes :
func SetupRoutes(apiBasePath string) {
	cbuHandler := http.HandlerFunc(handleCBUs)
	downloadHandler := http.HandlerFunc(handleDownload)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, cbusPath), middleware.Middleware(cbuHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, cbusPath), middleware.Middleware(downloadHandler))
}

func handleCBUs(w http.ResponseWriter, r *http.Request) {
	log.Println("invoking cbus")
	switch r.Method {
	case http.MethodGet:
		receiptList, err := GetCBUFileInfo()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(receiptList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		uploadFile(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 << 20)
	file, handler, err := r.FormFile("cbuFile")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileName := time.Now().Format("2006-01-02") + handler.Filename
	fileName = filepath.Join(CBUUploadPath, fileName)
	log.Println(fileName)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.Copy(f, file)
	w.WriteHeader(http.StatusCreated)
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", cbusPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fileName := urlPathSegments[1:][0]
	log.Println(urlPathSegments)
	file, err := os.Open(filepath.Join(CBUUploadPath, fileName))
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fHeader := make([]byte, 512)
	file.Read(fHeader)
	fContentType := http.DetectContentType(fHeader)

	stat, err := file.Stat()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fSize := strconv.FormatInt(stat.Size(), 10)
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", fContentType)
	w.Header().Set("Content-Length", fSize)
	file.Seek(0, 0)
	io.Copy(w, file)
}
