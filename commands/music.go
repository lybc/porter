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
        if matchSingle, err := regexp.MatchString("https://music.163.com/song/*", url); matchSingle && err == nil {
            return downloadSingle(url)
        } else if matchPlaylist, err := regexp.MatchString("https://music.163.com/playlist/*", url); matchPlaylist && err == nil {
            return downloadPlayList(url)
        } else if matchRadio, err := regexp.MatchString("https://music.163.com/radio/*", url); matchRadio && err == nil {
            return downloadRadio(url)
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

func downloadRadio(resourceUrl string) error {
    id := getIdByUrl(resourceUrl)
    radio := api.GetRadio(id)
    downloader := utils.NewDownloader("./", 3)
    for _, p := range radio.Programs {
        downloader.AppendResource(p.MainSong.GetFileName(), p.MainSong.GetStreamUrl())
    }

    downloader.Start()
    //fmt.Println(downloader.Resources)
    //wg := sync.WaitGroup{}
    //for _, p := range radio.Programs {
    //    wg.Add(1)
    //    go func(s api.Song, wg *sync.WaitGroup) {
    //        err := utils.DownloadFile(s.GetFileName(), s.GetStreamUrl(), nil)
    //        if err != nil {
    //            fmt.Println(err)
    //        }
    //        wg.Done()
    //    }(p.MainSong, &wg)
    //}
    //wg.Wait()
    return nil
}

func downloadPlayList(resourceUrl string) error {
    //u, err := url.Parse(resourceUrl)
    //if err != nil {
    //    panic(err)
    //}
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