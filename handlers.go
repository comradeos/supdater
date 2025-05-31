package main

import (
	"fmt"
	"net/http"
	"os"
)

func FileExists(w http.ResponseWriter, r *http.Request) {
	var db = DB
	var dataPath = DataPath

	println(db)
	println(dataPath)

	var fileName = r.URL.Query().Get("name")
	if fileName == "" {
		http.Error(w, "File name is required", http.StatusBadRequest)
		return
	}

	file, err := GetFileByName(db, fileName)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	if file == nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// получить ссылку на файл
	filePath := fmt.Sprintf("%s/%s", dataPath, file.Name)

	//fmt.Fprintf(w, "SHA256: %s\n", file.SHA256)

	// хочу что бы был доступ к файлу по ссылке http://192.168.0.101:8080/music.mp3

	fmt.Fprintf(w, "File found: %s\n", filePath)
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	var fileName = r.URL.Path[1:] // убираем первый слеш

	if fileName == "" {
		http.Error(w, "Filename required", http.StatusBadRequest)
		return
	}

	filePath := fmt.Sprintf("%s/%s", DataPath, fileName)

	// Проверка, существует ли файл
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))

	http.ServeFile(w, r, filePath)
}
