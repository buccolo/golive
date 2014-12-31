package main

import (
	"io/ioutil"
	"net/http"
	"runtime"
	"testing"
)

func TestMain(t *testing.T) {
	go Serve()

	response, err := http.Get("http://localhost:8108/golive_test.go")
	if err != nil {
		t.Errorf("We need port 8108 to run the test: %s", err)
		return
	}

	response_content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	_, filename, _, _ := runtime.Caller(0)
	file_content, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	if string(response_content) != string(file_content) {
		t.Errorf("The content via web and file-system are not the same.\nWeb:\n%s\nFile:\n%s", response_content, file_content)
	}
}
