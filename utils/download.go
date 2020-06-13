package utils

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "sync"
)

type Resource struct {
    Filename string
    Url string
}

type Downloader struct {
    wg *sync.WaitGroup
    pool chan *Resource
    HttpClient http.Client
    TargetDir string
    Resources []Resource
}

func NewDownloader(targetDir string, concurrent int) *Downloader {
    return &Downloader{
        wg: &sync.WaitGroup{},
        pool: make(chan *Resource, concurrent),
        TargetDir: targetDir,
    }
}

func (d *Downloader) AppendResource(filename, url string) {
    d.Resources = append(d.Resources, Resource{
        Filename: filename,
        Url:      url,
    })
}

func (d *Downloader) Download(resource Resource) error {
    defer d.wg.Done()
    d.pool <- &resource
    fmt.Println(resource.Filename, resource.Url)
    finalPath := d.TargetDir + "/" + resource.Filename
    target, err := os.Create(finalPath + ".tmp")

    if err != nil {
        return err
    }

    req, err := http.NewRequest(http.MethodGet, resource.Url, nil)
    if err != nil {
        return err
    }
    //if headers != nil {
    //    for k, v := range headers {
    //        req.Header.Add(k, v)
    //    }
    //}
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        target.Close()
        return err
    }
    defer resp.Body.Close()
    if _, err := io.Copy(target, resp.Body); err != nil {
        target.Close();
        return err
    }

    target.Close()
    if err := os.Rename(finalPath + ".tmp", finalPath); err != nil {
        return err
    }
    <- d.pool
    return nil
}

func (d *Downloader) Start() error {
    for _, resource := range d.Resources {
        d.wg.Add(1)
        go d.Download(resource)
    }
    d.wg.Wait()
    return nil
}
