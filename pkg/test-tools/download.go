package tools

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const testDir = "image/tiff/tests/raw"

func DownloadRAW(link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	destFile := testDir + "/" + u.Path
	destFolder := filepath.Dir(destFile)
	err = os.MkdirAll(destFolder, 0777)
	if err != nil {
		return "", err
	}
	_, err = os.Stat(destFile)
	if os.IsNotExist(err) {
		resp, err := http.Get(link)
		if err != nil {
			return "", err
		}

		defer resp.Body.Close()

		fd, err := os.Create(destFile)
		if err != nil {
			return "", err
		}

		_, err = io.Copy(fd, resp.Body)
		if err != nil {
			return "", err
		}

		fmt.Println("downloaded", link, "to", destFile)
	} else {
		if err != nil {
			return "", err
		}
		fmt.Println("Skipping existing raw file", link)
	}
	return destFile, nil
}

func DownloadTestRAW(t testing.TB, link string) string {
	t.Helper()
	dest, err := DownloadRAW(link)
	require.NoError(t, err)
	return dest
}
