// Fetch downloads the URL and returns the
// name and length of the local file.

package main

import (
	"os"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URl.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but preferably an error from copy
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
