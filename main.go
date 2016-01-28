package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//AppInfo contains info about this app
type AppInfo struct {
	Version   string    `json:"version"`
	StartedAt time.Time `json:"startedAt"`
}

var appInfo AppInfo

func init() {
	appInfo = AppInfo{Version: "0.1.0.0", StartedAt: time.Now()}
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(appInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func getRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/info", infoHandler)
	return r
}

func main() {
	http.ListenAndServe(":8080", getRouter())
}
