package api

import (
	"encoding/json"
	"fmt"
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

type Song struct {
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
}

type Songs struct {
	Songs      []Song   `json:"songs"`
	Equalizers struct{} `json:"equalizers"`
	Code       int      `json:"code"`
}

type RadioPayload struct {
	Count    int `json:"count"`
	Code     int `json:"code"`
	Programs []struct {
		MainSong Song   `json:"mainSong"`
		Songs    string `json:"songs"`
		Dj       struct {
			DefaultAvatar      bool        `json:"defaultAvatar"`
			Province           int         `json:"province"`
			AuthStatus         int         `json:"authStatus"`
			Followed           bool        `json:"followed"`
			AvatarURL          string      `json:"avatarUrl"`
			AccountStatus      int         `json:"accountStatus"`
			Gender             int         `json:"gender"`
			City               int         `json:"city"`
			Birthday           int64       `json:"birthday"`
			UserID             int         `json:"userId"`
			UserType           int         `json:"userType"`
			Nickname           string      `json:"nickname"`
			Signature          string      `json:"signature"`
			Description        string      `json:"description"`
			DetailDescription  string      `json:"detailDescription"`
			AvatarImgID        int64       `json:"avatarImgId"`
			BackgroundImgID    int64       `json:"backgroundImgId"`
			BackgroundURL      string      `json:"backgroundUrl"`
			Authority          int         `json:"authority"`
			Mutual             bool        `json:"mutual"`
			ExpertTags         interface{} `json:"expertTags"`
			Experts            interface{} `json:"experts"`
			DjStatus           int         `json:"djStatus"`
			VipType            int         `json:"vipType"`
			RemarkName         interface{} `json:"remarkName"`
			AvatarImgIDStr     string      `json:"avatarImgIdStr"`
			BackgroundImgIDStr string      `json:"backgroundImgIdStr"`
			AvatarImgIDStr2    string      `json:"avatarImgId_str"`
			Brand              string      `json:"brand"`
		}
		BlurCoverUrl string `json:"blueCoverUrl"`
		Radio        struct {
			Dj                    interface{} `json:"dj"`
			Category              string      `json:"category"`
			Buyed                 bool        `json:"buyed"`
			Price                 int         `json:"price"`
			OriginalPrice         int         `json:"originalPrice"`
			DiscountPrice         interface{} `json:"discountPrice"`
			PurchaseCount         int         `json:"purchaseCount"`
			LastProgramName       interface{} `json:"lastProgramName"`
			Videos                interface{} `json:"videos"`
			Finished              bool        `json:"finished"`
			UnderShelf            bool        `json:"underShelf"`
			LiveInfo              interface{} `json:"liveInfo"`
			CategoryID            int         `json:"categoryId"`
			CreateTime            int64       `json:"createTime"`
			LastProgramID         int         `json:"lastProgramId"`
			FeeScope              int         `json:"feeScope"`
			PicID                 int64       `json:"picId"`
			ProgramCount          int         `json:"programCount"`
			SubCount              int         `json:"subCount"`
			LastProgramCreateTime int64       `json:"lastProgramCreateTime"`
			RadioFeeType          int         `json:"radioFeeType"`
			PicURL                string      `json:"picUrl"`
			Desc                  string      `json:"desc"`
			Name                  string      `json:"name"`
			ID                    int         `json:"id"`
		}
		Alg                      interface{}   `json:"alg"`
		AuditStatus              int64         `json:"auditStatus"`
		BdAuditStatus            int64         `json:"bdAuditStatus"`
		Buyed                    bool          `json:"buyed"`
		CanReward                bool          `json:"canReward"`
		Channels                 []interface{} `json:"channels"`
		CommentCount             int64         `json:"commentCount"`
		CommentThreadID          string        `json:"commentThreadId"`
		CoverURL                 string        `json:"coverUrl"`
		CreateTime               int64         `json:"createTime"`
		Description              string        `json:"description"`
		Duration                 int64         `json:"duration"`
		FeeScope                 int64         `json:"feeScope"`
		H5Links                  interface{}   `json:"h5Links"`
		ID                       int64         `json:"id"`
		IsPublish                bool          `json:"isPublish"`
		LikedCount               int64         `json:"likedCount"`
		ListenerCount            int64         `json:"listenerCount"`
		LiveInfo                 interface{}   `json:"liveInfo"`
		MainTrackID              int64         `json:"mainTrackId"`
		Name                     string        `json:"name"`
		ProgramDesc              interface{}   `json:"programDesc"`
		ProgramFeeType           int64         `json:"programFeeType"`
		PubStatus                int64         `json:"pubStatus"`
		Reward                   bool          `json:"reward"`
		Score                    int64         `json:"score"`
		SerialNum                int64         `json:"serialNum"`
		ShareCount               int64         `json:"shareCount"`
		SmallLanguageAuditStatus int64         `json:"smallLanguageAuditStatus"`
		Subscribed               bool          `json:"subscribed"`
		SubscribedCount          int64         `json:"subscribedCount"`
		TitbitImages             interface{}   `json:"titbitImages"`
		Titbits                  interface{}   `json:"titbits"`
		TrackCount               int64         `json:"trackCount"`
		VideoInfo                interface{}   `json:"videoInfo"`
	}
}

func (s Song) GetStreamUrl() string {
	return fmt.Sprintf("http://music.163.com/song/media/outer/url?id=%d.mp3", s.ID)
}

func (s Song) GetFileName() string {
	return fmt.Sprintf("%s.mp3", s.Name)
}

func GetSongsInfo(ids []string) Songs {
	var s = Songs{}
	tmp, _ := json.Marshal(ids)
	NewClient().Get("http://music.163.com/api/song/detail/?ids="+string(tmp), &s)
	return s
}

func GetRadio(id string) RadioPayload {
	var payload = RadioPayload{}
	url := fmt.Sprintf("http://music.163.com/api/dj/program/byradio/?radioId=%s&ids=[%s]&csrf_token=", id, id)
	NewClient().Get(url, &payload)
	return payload
}
