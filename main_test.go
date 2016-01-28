package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	server  *httptest.Server
	infoURL string
)

func init() {
	server = httptest.NewServer(getRouter())
	infoURL = fmt.Sprintf("%v/%v", server.URL, "info")
}

func TestInfoEndpoint(t *testing.T) {
	request, err := http.NewRequest("GET", infoURL, nil)
	if err != nil {
		t.Error(err)
	}
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %v, got %v", http.StatusOK, res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	var infoResponse AppInfo
	if err := json.Unmarshal(body, &infoResponse); err != nil {
		t.Error(err)
	}

	if infoResponse.Version != "0.1.0.0" {
		t.Errorf("Expected version %v, got %v", "0.1.0.0", infoResponse.Version)
	}
}
