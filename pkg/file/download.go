package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Download(url, path, fileName string) (fullname string, err error) {
	fullname = filepath.Join(path, fileName)
	out, err := os.Create(fullname)
	if err != nil {
		return "", fmt.Errorf("create file err: %w", err)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("http get err: %w", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("io copy err: %w", err)
	}

	return fullname, nil
}
