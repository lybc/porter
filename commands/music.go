package commands

import (
    "fmt"
    "github.com/urfave/cli"
    "net/url"
    "porter/api"
    "regexp"
)

const (
    MusicMediaUrl = "http://music.163.com/song/media/outer/url?id=%s.mp3"
)

var musicCmd = cli.Command{
    Name:  "music",
    Usage: "下载音乐",
    Flags: []cli.Flag{
        &cli.StringFlag{
            Name:  "output",
            Usage: "下载文件输出路径",
            Value: "./",
        },
    },
    Action: func(c *cli.Context) error {
        url := c.Args().Get(0)
        if matchSingle, err := regexp.MatchString("https://music.163.com/song/*", url); matchSingle && err == nil {
            return downloadSingle(url)
        } else if matchPlaylist, err := regexp.MatchString("https://music.163.com/playlist/*", url); matchPlaylist && err == nil {
            return downloadPlayList(url)
        } else if matchRadio, err := regexp.MatchString("https://music.163.com/radio/*", url); matchRadio && err == nil {
            return downloadRadio(url, c)
        }
        return nil
    },
}

func init() {
    RootCmd.Commands = append(RootCmd.Commands, musicCmd)
}

func getIdByUrl(resourceUrl string) string {
    if u, err := url.Parse(resourceUrl); err == nil {
        return u.Query().Get("id")
    }
    return ""
}

func downloadRadio(resourceUrl string, c *cli.Context) error {
    //id := getIdByUrl(resourceUrl)
    //radio := api.GetRadio(id)
    //fmt.Println(c.String("output"))
    //downloader := utils.NewDownloader(c.String("output"))
    //for _, p := range radio.Programs {
    //    downloader.AppendResource(p.MainSong.GetFileName(), p.MainSong.GetStreamUrl())
    //}
    //downloader.Start()
    return nil
}

func downloadPlayList(resourceUrl string) error {
    u, err := url.Parse(resourceUrl)
    if err != nil {
        return err
    }
    playlistId := u.Query().Get("id")
    api := api.Netease{}
    playList := api.GetPlayListDetail(playlistId)
    if playList.Code != 200 {
        return fmt.Errorf("获取网易云歌单详情失败")
    }
    fmt.Println(playList.Playlist.TrackCount)

    for _, song := range playList.Playlist.Tracks {
        fmt.Println(song.Name)
    }
    return nil
}

func downloadSingle(resourceUrl string) error {
    //u, err := url.Parse(resourceUrl)
    //if err != nil {
    //    panic(err)
    //}
    //musicId := u.Query().Get("id")
    //songs := api.GetSongsInfo([]string{musicId})
    //
    //if songs.Code != 200 {
    //    return fmt.Errorf("获取歌曲信息失败")
    //}
    //
    //for _, song := range songs.Songs {
    //    name := fmt.Sprintf("%s.mp3", song.Name)
    //    mp3Url := fmt.Sprintf(MusicMediaUrl, musicId)
    //    utils.DownloadFile(name, mp3Url, nil)
    //}

    return nil
}
