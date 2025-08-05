package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const uploadDir = "./uploads/"

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method %s not supported.", r.Method)
		return
	}

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Println("form parsing error:", err)
		http.Error(w, "form reading error", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println("error receiving file:", err)
		http.Error(w, "file not found", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("error reading file:", err)
		http.Error(w, "error reading file", http.StatusInternalServerError)
		return
	}

	convertedData := service.ConvertMorseCode(string(data))

	timestamp := time.Now().UTC().Format("2006-01-02T15_04_05.000000000Z")
	filename := fmt.Sprintf("%v%s", timestamp, filepath.Ext(handler.Filename))
	savePath := filepath.Join(uploadDir, filename)

	f, err := os.Create(savePath)
	if err != nil {
		log.Println("error creating file:", err)
		http.Error(w, "file write error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = f.Write([]byte(convertedData))
	if err != nil {
		log.Println("error writing to file:", err)
		http.Error(w, "file write error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, convertedData)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "file transfer error", http.StatusInternalServerError)
		return
	}
}
