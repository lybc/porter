package commands

import (
    "fmt"
    "github.com/urfave/cli"
    "net/url"
    "porter/api"
    "porter/utils"
    "regexp"
)

const (
    MusicMediaUrl = "http://music.163.com/song/media/outer/url?id=%s.mp3"
)

var musicCmd = cli.Command{
    Name: "music",
    Usage: "下载音乐",
    Action: func(c *cli.Context) error {
        url := c.Args().Get(0)
        matchSingle, _ := regexp.MatchString("https://music.163.com/song/*", url)
        if matchSingle {
            return downloadSingle(url)
        }
        //https://music.163.com/playlist?id=38196761&userid=44216499
        matchPlaylist, _ := regexp.MatchString("https://music.163.com/playlist/*", url)
        if matchPlaylist {
            return downloadPlayList(url)
        }
        return nil
    },
}

func init() {
    RootCmd.Commands = append(RootCmd.Commands, musicCmd)
}

func downloadPlayList(resourceUrl string) error {
    u, err := url.Parse(resourceUrl)
    if err != nil {
        panic(err)
    }
    //playlistId := u.Query().Get("id")

    return nil
}

func downloadSingle(resourceUrl string) error {
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