package api

import (
    "encoding/json"
)

type Artist struct {
	Name        string        `json:"name"`
	ID          int           `json:"id"`
	PicID       int           `json:"picId"`
	Img1V1ID    int           `json:"img1v1Id"`
	BriefDesc   string        `json:"briefDesc"`
	PicURL      string        `json:"picUrl"`
	Img1V1URL   string        `json:"img1v1Url"`
	AlbumSize   int           `json:"albumSize"`
	Alias       []interface{} `json:"alias"`
	Trans       string        `json:"trans"`
	MusicSize   int           `json:"musicSize"`
	TopicPerson int           `json:"topicPerson"`
}

type Album struct {
	Name            string        `json:"name"`
	ID              int           `json:"id"`
	Type            string        `json:"type"`
	Size            int           `json:"size"`
	PicID           int64         `json:"picId"`
	BlurPicURL      string        `json:"blurPicUrl"`
	CompanyID       int           `json:"companyId"`
	Pic             int64         `json:"pic"`
	PicURL          string        `json:"picUrl"`
	PublishTime     int64         `json:"publishTime"`
	Description     string        `json:"description"`
	Tags            string        `json:"tags"`
	Company         string        `json:"company"`
	BriefDesc       string        `json:"briefDesc"`
	Artist          Artist        `json:"artist"`
	Songs           []interface{} `json:"songs"`
	Alias           []interface{} `json:"alias"`
	Status          int           `json:"status"`
	CopyrightID     int           `json:"copyrightId"`
	CommentThreadID string        `json:"commentThreadId"`
	Artists         []Artist      `json:"artists"`
	SubType         string        `json:"subType"`
	TransName       interface{}   `json:"transName"`
	Mark            int           `json:"mark"`
	PicIDStr        string        `json:"picId_str"`
}

type Music struct {
	Name        interface{} `json:"name"`
	ID          int         `json:"id"`
	Size        int         `json:"size"`
	Extension   string      `json:"extension"`
	Sr          int         `json:"sr"`
	DfsID       int         `json:"dfsId"`
	Bitrate     int         `json:"bitrate"`
	PlayTime    int         `json:"playTime"`
	VolumeDelta int         `json:"volumeDelta"`
}

type Songs struct {
	Songs []struct {
		Name            string        `json:"name"`
		ID              int           `json:"id"`
		Position        int           `json:"position"`
		Alias           []interface{} `json:"alias"`
		Status          int           `json:"status"`
		Fee             int           `json:"fee"`
		CopyrightID     int           `json:"copyrightId"`
		Disc            string        `json:"disc"`
		No              int           `json:"no"`
		Artists         Artist        `json:"artists"`
		Album           Album         `json:"album"`
		Starred         bool          `json:"starred"`
		Popularity      float64       `json:"popularity"`
		Score           int           `json:"score"`
		StarredNum      int           `json:"starredNum"`
		Duration        int           `json:"duration"`
		PlayedNum       int           `json:"playedNum"`
		DayPlays        int           `json:"dayPlays"`
		HearTime        int           `json:"hearTime"`
		Ringtone        interface{}   `json:"ringtone"`
		Crbt            interface{}   `json:"crbt"`
		Audition        interface{}   `json:"audition"`
		CopyFrom        string        `json:"copyFrom"`
		CommentThreadID string        `json:"commentThreadId"`
		RtURL           interface{}   `json:"rtUrl"`
		Ftype           int           `json:"ftype"`
		RtUrls          []interface{} `json:"rtUrls"`
		Copyright       int           `json:"copyright"`
		TransName       interface{}   `json:"transName"`
		Sign            interface{}   `json:"sign"`
		Mark            int           `json:"mark"`
		NoCopyrightRcmd interface{}   `json:"noCopyrightRcmd"`
		HMusic          Music         `json:"hMusic"`
		MMusic          Music         `json:"mMusic"`
		LMusic          Music         `json:"lMusic"`
		BMusic          Music         `json:"bMusic"`
		Mvid            int           `json:"mvid"`
		Mp3URL          interface{}   `json:"mp3Url"`
		Rtype           int           `json:"rtype"`
		Rurl            interface{}   `json:"rurl"`
	} `json:"songs"`
	Equalizers struct {
	} `json:"equalizers"`
	Code int `json:"code"`
}

func GetSongsInfo(ids []string) Songs {
    var s = Songs{}
    tmp, _ := json.Marshal(ids)
    NewClient().Get("http://music.163.com/api/song/detail/?ids=" + string(tmp), &s)
    return s
}
