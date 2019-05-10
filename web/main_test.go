package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	path, err := BuildBinary()
	if err != nil {
		panic(err)
	}

	go RunBinary(path)
	time.Sleep(2 * time.Second)
	status := m.Run()

	os.Exit(status)
}

func BuildBinary() (string, error) {
	output, err := exec.Command("go", "build", "-o", "/tmp/server", "/Users/alex/go/src/github.com/burov/courses/web/02_http_crud/").CombinedOutput()
	fmt.Println(string(output))

	if err != nil {
		return "", err
	}

	return "/tmp/server", nil

}

func RunBinary(path string) error {
	output, err := exec.Command("/tmp/server").Output()
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}
