package drivers

import (
    "fmt"
    "github.com/bitly/go-simplejson"
    "net/http"
    "net/url"
)

const (
    MusicInfoUrl  = "http://music.163.com/api/song/detail/?id=%s&ids=[%s]&csrf_token="
    MusicMediaUrl = "http://music.163.com/song/media/outer/url?id=%s.mp3"
)

// 下载网易云音乐
type NeteaseMusic struct {
    ID     string
    Name   string
    Singer string
}

func (n *NeteaseMusic) Download(resourceUrl string, outputPath string) bool {
    mp3Url := fmt.Sprintf(MusicMediaUrl, n.ID)
    fmt.Println("正在下载：", n.Name, n.Singer)
    destPath := outputPath + "/" + fmt.Sprintf("%s(%s).mp3", n.Name, n.Singer)
    httpDownload(mp3Url, destPath)
    return true
}

func NewNeteaseMusic(fromUrl string) *NeteaseMusic {
    u, err := url.Parse(fromUrl)
    if err != nil {
        panic(err)
    }
    musicId := u.Query().Get("id")
    response, _ := http.Get(fmt.Sprintf(MusicInfoUrl, musicId, musicId))
    if response.StatusCode != http.StatusOK {
        panic("无法获取到歌曲信息")
    }

    defer response.Body.Close()
    musicInfo, _ := simplejson.NewFromReader(response.Body)
    musicInfo = musicInfo.Get("songs").GetIndex(0)
    return &NeteaseMusic{
        ID: musicId,
        Name: musicInfo.Get("name").MustString(),
        Singer: musicInfo.Get("artists").GetIndex(0).Get("name").MustString(),
    }
}