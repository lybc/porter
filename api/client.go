package api

import (
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
)

func NewClient() *Client {
    tr := http.DefaultTransport
    http := &http.Client{Transport: tr}
    client := &Client{http: http}
    return client
}

type Client struct {
    http *http.Client
}

// 下载文件
func (c Client) Download(url string, target string) error {
    file, err := os.Create(target)
    if err != nil {
        return err
    }

    defer file.Close()

    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    io.Copy(file, resp.Body)
    return nil
}

func (c Client) Get(url string, data interface{}) error {
    req, err := http.NewRequest(http.MethodGet, url, nil)
    
    if err != nil {
        return err
    }

    resp, err := c.http.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    success := resp.StatusCode >= 200 && resp.StatusCode < 300

    if !success {
        return fmt.Errorf("http error, '%s' failed (%d)", resp.Request.URL, resp.StatusCode)
    }

    b, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    err = json.Unmarshal(b, &data)
    if err != nil {
        return err
    }

    return nil
}