package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_bank(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "306680") {
		t.Errorf("wrong balance")
	}

}
//output:
// go test bank\main_test.go bank\main.go
