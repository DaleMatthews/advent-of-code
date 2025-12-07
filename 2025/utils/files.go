package utils

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func ReadInput(t *testing.T, filename string, options ...bool) string {
	trim := true
	if len(options) > 0 {
		trim = options[0]
	}
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		t.Fatal("Failed to get caller information. Can't determine which day to load.")
	}

	filePath := fmt.Sprintf("%s/%s", filepath.Dir(file), filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatal("Input file not found. Make sure to run `make init` to complete the one-time setup.")
	}

	if trim {
		return string(bytes.TrimSpace(content))
	}
	return strings.ReplaceAll(string(content), "\r\n", "\n")
	// return string(content)
}
