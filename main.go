package main

import (
	"log"
	"mastergo-font-linux/internal/dao"
	"mastergo-font-linux/internal/middleware"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/local-fonts", middleware.CORS(http.HandlerFunc(dao.GetLocalFontsHandler)))
	mux.Handle("/font-file", middleware.CORS(http.HandlerFunc(dao.GetFontFileHandler)))
	mux.Handle("/ziyou-fonts", middleware.CORS(http.HandlerFunc(dao.GetZiYouFontsHandler)))
	mux.Handle("/cache-fonts", middleware.CORS(http.HandlerFunc(dao.GetCacheFontsHandler)))
	mux.Handle("/upload-font", middleware.CORS(http.HandlerFunc(dao.UploadFontHandler)))
	log.Println("Starting server on :26062")
	err := http.ListenAndServe(":26062", mux)
	log.Fatal(err)
}
