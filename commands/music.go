package commands

import (
    "fmt"
    "github.com/bitly/go-simplejson"
    "github.com/urfave/cli"
    "net/http"
    "net/url"
)

const (
    MusicInfoUrl  = "http://music.163.com/api/song/detail/?id=%s&ids=[%s]&csrf_token="
    MusicMediaUrl = "http://music.163.com/song/media/outer/url?id=%s.mp3"
)

var musicCmd = cli.Command{
    Name: "music",
    Usage: "下载音乐",
    Action: func(c *cli.Context) error {
        url := c.Args().Get(0)
        output := c.Args().Get(1)
        downloadNetease(url, output)
        return nil
    },
}

func init() {
    RootCmd.Commands = append(RootCmd.Commands, musicCmd)
}

func downloadNetease(resourceUrl string, outputPath string) bool {
    u, err := url.Parse(resourceUrl)
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
    mp3Url := fmt.Sprintf(MusicMediaUrl, musicId)
    name := musicInfo.Get("name").MustString()
    singer := musicInfo.Get("artists").GetIndex(0).Get("name").MustString()
    fmt.Println("正在下载：", name, singer)

    destPath := outputPath + "/" + fmt.Sprintf("%s(%s).mp3", name, singer)
    httpDownload(mp3Url, destPath)
    return true
}