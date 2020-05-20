package commands

import (
    "io"
    "net/http"
    "os"
)

type downloader interface {
    Download(resourceUrl string, outputPath string) bool
}

func httpDownload(url, output string) {
    file, err := os.Create(output)
    if err != nil {
        panic(err)
    }

    defer file.Close()

    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    io.Copy(file, resp.Body)
}
