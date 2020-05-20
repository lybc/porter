package api

import (
    "fmt"
)

type PlayListResult struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    TTL     int    `json:"ttl"`
    Data    []struct {
        Cid       int    `json:"cid"`
        Page      int    `json:"page"`
        From      string `json:"from"`
        Part      string `json:"part"`
        Duration  int    `json:"duration"`
        Vid       string `json:"vid"`
        Weblink   string `json:"weblink"`
        Dimension struct {
            Width  int `json:"width"`
            Height int `json:"height"`
            Rotate int `json:"rotate"`
        } `json:"dimension"`
    } `json:"data"`
}

type PlayUrlResult struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    TTL     int    `json:"ttl"`
    Data    struct {
        From              string   `json:"from"`
        Result            string   `json:"result"`
        Message           string   `json:"message"`
        Quality           int      `json:"quality"`
        Format            string   `json:"format"`
        Timelength        int      `json:"timelength"`
        AcceptFormat      string   `json:"accept_format"`
        AcceptDescription []string `json:"accept_description"`
        AcceptQuality     []int    `json:"accept_quality"`
        VideoCodecid      int      `json:"video_codecid"`
        SeekParam         string   `json:"seek_param"`
        SeekType          string   `json:"seek_type"`
        Durl              []struct {
            Order     int      `json:"order"`
            Length    int      `json:"length"`
            Size      int      `json:"size"`
            Ahead     string   `json:"ahead"`
            Vhead     string   `json:"vhead"`
            URL       string   `json:"url"`
            BackupURL []string `json:"backup_url"`
        } `json:"durl"`
    } `json:"data"`
}

// 获取分P播放的结果
func GetPlayList(bvid string) PlayListResult {
    url := fmt.Sprintf("https://api.bilibili.com/x/player/pagelist?bvid=%s&jsonp=jsonp", bvid)
    var playList = PlayListResult{}
    NewClient().Get(url, &playList)
    return playList
}

func GetPlayUrl(bvid string, cid int) PlayUrlResult {
    url := fmt.Sprintf("http://api.bilibili.com/x/player/playurl?bvid=%s&cid=%d", bvid, cid)
    var playUrl = PlayUrlResult{}
    NewClient().Get(url, &playUrl)
    return playUrl
}


