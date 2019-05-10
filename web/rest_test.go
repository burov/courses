package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestEmptyList(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/users/")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()

	content, _ := ioutil.ReadAll(resp.Body)
	if string(content) != "[]" {
		t.Fatalf("wrong output, expected %s, got %s", "{}", string(content))
	}
}
