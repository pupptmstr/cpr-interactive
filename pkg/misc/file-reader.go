package misc

import (
	"fmt"
	"os"
)

func ReadAssetAsString(path string) (string, error) {
	content, err := ReadAssetAsBytes(path)
	return string(content), err
}

func ReadAssetAsBytes(path string) ([]byte, error) {
	fullPath := fmt.Sprintf("assets/%s", path)
	content, err := os.ReadFile(fullPath)
	return content, err
}
