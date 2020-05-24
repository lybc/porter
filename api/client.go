package api

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
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