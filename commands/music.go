package commands

import (
    "fmt"
    "github.com/urfave/cli"
    "net/url"
    "porter/api"
    "porter/utils"
)

const (
    MusicMediaUrl = "http://music.163.com/song/media/outer/url?id=%s.mp3"
)

var musicCmd = cli.Command{
    Name: "music",
    Usage: "下载音乐",
    Action: func(c *cli.Context) error {
        url := c.Args().Get(0)
        downloadNetease(url)
        return nil
    },
}

func init() {
    RootCmd.Commands = append(RootCmd.Commands, musicCmd)
}

func downloadNetease(resourceUrl string) error {
    u, err := url.Parse(resourceUrl)
    if err != nil {
        panic(err)
    }
    musicId := u.Query().Get("id")
    songs := api.GetSongsInfo([]string{musicId})

    if songs.Code != 200 {
        return fmt.Errorf("获取歌曲信息失败")
    }

    for _, song := range songs.Songs {
        name := fmt.Sprintf("%s.mp3", song.Name)
        mp3Url := fmt.Sprintf(MusicMediaUrl, musicId)
        utils.DownloadFile(name, mp3Url, nil)
    }

    return nil
}