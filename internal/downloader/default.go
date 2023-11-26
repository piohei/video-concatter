package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// defaultDownloader is a default implementation of Downloader that uses in-process go code.
type defaultDownloader struct {
	retries int
}

// Download downloads file from the passed url and writes it to target path. On error retries configured number of times.
// In case of error returns the last error that occurred.
func (d *defaultDownloader) Download(url, target string) error {
	var err error
	for i := 0; i <= d.retries; i++ {
		if err = d.downloadFile(url, target); err == nil {
			break
		}
	}
	return err
}

// setRetries configures number of retries.
func (d *defaultDownloader) setRetries(n int) {
	d.retries = n
}

// downloadFile downloads file specified by url to a target path.
func (d *defaultDownloader) downloadFile(url, target string) error {
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err = out.Close()
		log.Printf("error closing output file for download: %s", err)
	}(out)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		// [TODO] Error here could be logged. Not covered in PoC.
		err = Body.Close()
		log.Printf("error closing request body: %s", err)
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

var _ Downloader = (*defaultDownloader)(nil)
