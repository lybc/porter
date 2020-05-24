package commands

import (
    "fmt"
    "github.com/urfave/cli"
    "net/url"
    "porter/api"
    "porter/utils"
    "strings"
)

var videoCmd = cli.Command{
   Name:  "video",
   Usage: "下载视频",
    Action: func(c *cli.Context) error {
        downloadBilibili(c)
        return nil
    },
}

func init() {
   RootCmd.Commands = append(RootCmd.Commands, videoCmd)
}

// 解析URL中的bvid
func getBvid(resourceUrl string) (string, error) {
    u, err := url.Parse(resourceUrl)
    if err != nil {
        return "", fmt.Errorf("Invalid url: %s", resourceUrl)
    }

    path := strings.Split(u.Path, "/")
    return path[len(path) - 1], nil
}

func downloadBilibili(c *cli.Context) error {
    url := c.Args().Get(0)
    bvid, err := getBvid(url)
    if err != nil {
        return fmt.Errorf("解析bvid失败：%s", url)
    }

    playList := api.GetPlayList(bvid)
    if playList.Code != 0 {
        return fmt.Errorf("获取播放列表接口请求失败：%s", playList.Message)
    }

    cid := playList.Data[0].Cid
    playUrl := api.GetPlayUrl(bvid, cid)
    if playUrl.Code != 0 {
        return fmt.Errorf("获取视频流接口请求失败：%s", playUrl.Message)
    }
    durl := playUrl.Data.Durl[0]

    headers := map[string]string{
        "Accept": "*/*",
        "Accept-Language": "en-US,en;q=0.5",
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36",
        "Referer": url,
    }

    utils.DownloadFile(c.Args().Get(1), durl.URL, headers)
    return nil
}
