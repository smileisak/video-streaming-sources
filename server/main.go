package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexPage).Methods("GET")
	router.HandleFunc("/media/{mID:[0-9]+}/stream/1080p.m3u8", streamHandler).Methods("GET")
	//router.HandleFunc("/media/{mID:[0-9]+}/stream2/{segName:[0-9a-z]+.ts}", streamHandler).Methods("GET")
	router.HandleFunc("/media/{mID:[0-9]+}/stream/{segName:[0-9a-z]+.ts}", streamHandler).Methods("GET")
	return router
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func streamHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	mID, err := strconv.Atoi(vars["mID"])
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	log.Printf("streamHandler called for id: %d.", mID)

	segName, ok := vars["segName"]
	if !ok {
		mediaBase := getMediaBase(mID)
		m3u8Name := "1080p.m3u8"
		serveHlsM3u8(response, request, mediaBase, m3u8Name)
	} else {
		mediaBase := getMediaBase(mID)
		serveHlsTs(response, request, mediaBase, segName)
	}
}

func getMediaBase(mID int) string {
	mediaRoot := "bunny_1080p_60fps"
	return fmt.Sprintf("%s/%d", mediaRoot, mID)
}

func serveHlsM3u8(w http.ResponseWriter, r *http.Request, mediaBase, m3u8Name string) {
	mediaFile := fmt.Sprintf("%s/hls/%s", mediaBase, m3u8Name)
	http.ServeFile(w, r, mediaFile)
	//w.Header().Set("Content-Type", "application/x-mpegURL")
	w.Header().Set("Content-Type", "audio/mpegurl")
}

func serveHlsTs(w http.ResponseWriter, r *http.Request, mediaBase, segName string) {
	mediaFile := fmt.Sprintf("%s/hls/%s", mediaBase, segName)
	http.ServeFile(w, r, mediaFile)
	//w.Header().Set("Content-Type", "video/MP2T")
	w.Header().Set("Content-Type", "application/octet-stream")
}

func main() {
	log.Print("Listening on port :8000")

	http.Handle("/", handlers())
	http.ListenAndServe(":8000", nil)

}
