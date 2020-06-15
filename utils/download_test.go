package utils

import (
    "os"
    "testing"
)

func TestDownload(t *testing.T)  {
    links := make(map[string]string)
    links["001.jpg"] = "http://222.186.12.239:10010/ksacf_20190731/001.jpg"
    links["002.jpg"] = "http://222.186.12.239:10010/ksacf_20190731/002.jpg"
    links["003.jpg"] = "http://222.186.12.239:10010/ksacf_20190731/003.jpg"
    links["004.jpg"] = "http://222.186.12.239:10010/ksacf_20190731/004.jpg"
    links["005.jpg"] = "http://222.186.12.239:10010/ksacf_20190731/005.jpg"
    links["006.jpg"] = "http://222.186.12.239:10010/ksacf_20190731/006.jpg"
    downloader := NewDownloader("./")
    for k, v := range links {
        os.Remove("./" + k)
        downloader.AppendResource(k, v)
    }
    downloader.Start()

    for k, _ := range links {
        if !IsFile("./" + k) {
            t.Fail()
        }
        os.Remove("./" + k)
    }
}