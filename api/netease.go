package api

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

const (
	NONCE   = "0CoJUm6Qyw8W8jud"
	IV      = "0102030405060708"
	KEYS    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	PUB_KEY = "010001"
	MODULUS = "00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
)

func addRSAPadding(encText string, modulus string) string {
	ml := len(modulus)
	for i := 0; ml > 0 && modulus[i:i+1] == "0"; i++ {
		ml--
	}
	num := ml - len(encText)
	prefix := ""
	for i := 0; i < num; i++ {
		prefix += "0"
	}
	return prefix + encText
}

// 网易云接口加密
type ParamsCrypto struct {
	Plaint string   // 明文
	Cipher struct { // 密文
		Param     string
		EncSecKey string
	}
}

// rsa加密
func (c *ParamsCrypto) rsaEncrypt(secKey string, pubKey string, modulus string) string {
	// 倒序 key
	rKey := ""
	for i := len(secKey) - 1; i >= 0; i-- {
		rKey += secKey[i : i+1]
	}
	// 将 key 转 ascii 编码 然后转成 16 进制字符串
	hexRKey := ""
	for _, char := range []rune(rKey) {
		hexRKey += fmt.Sprintf("%x", int(char))
	}
	// 将 16进制 的 三个参数 转为10进制的 bigint
	bigRKey, _ := big.NewInt(0).SetString(hexRKey, 16)
	bigPubKey, _ := big.NewInt(0).SetString(pubKey, 16)
	bigModulus, _ := big.NewInt(0).SetString(modulus, 16)
	// 执行幂乘取模运算得到最终的bigint结果
	bigRs := bigRKey.Exp(bigRKey, bigPubKey, bigModulus)
	// 将结果转为 16进制字符串
	hexRs := fmt.Sprintf("%x", bigRs)
	// 可能存在不满256位的情况，要在前面补0补满256位
	return addRSAPadding(hexRs, modulus)
}

// 创建一个随机字符串作为secret key
func (c *ParamsCrypto) createSecretKey(size int) string {
	random := make([]byte, size)
	for i := range random {
		random[i] = KEYS[rand.Intn(len(KEYS))]
	}
	return string(random)
}

func (c *ParamsCrypto) Encrypt() {
	sk := c.createSecretKey(16)
	c.Cipher.Param = c.aesEncrypt(c.aesEncrypt(c.Plaint, NONCE), sk)
	c.Cipher.EncSecKey = c.rsaEncrypt(sk, PUB_KEY, MODULUS)
}

func (c *ParamsCrypto) aesEncrypt(v string, sk string) string {
	block, err := aes.NewCipher([]byte(sk))
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	valueBytes := []byte(v)
	blockSize := block.BlockSize()
	padding := blockSize - len(valueBytes)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	valueBytes = append(valueBytes, paddingText...)

	blockMode := cipher.NewCBCEncrypter(block, []byte(IV))
	cipherText := make([]byte, len(valueBytes))
	blockMode.CryptBlocks(cipherText, valueBytes)

	return base64.StdEncoding.EncodeToString(cipherText)
}

type Netease struct{}

// 发送HTTP请求到网易云服务器
func (n *Netease) request(httpMethod string, url string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Referer", "http://music.163.com")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4")
	request.Header.Set("Cookie", "appver=2.0.2")
	// todo 随机UA
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Fatal("http 请求错误")
		return nil, fmt.Errorf("http request error, get: %d", resp.StatusCode)
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

// 获取歌单详情接口
func (n *Netease) GetPlayListDetail(id string) *PlayListPayload {
	c := ParamsCrypto{Plaint: fmt.Sprintf(`{"id": %s, "s": "8", "n": 10000, "csrf_token": ""}`, id)}
	c.Encrypt()
	params := url.Values{}
	params.Set("params", c.Cipher.Param)
	params.Set("encSecKey", c.Cipher.EncSecKey)
	responseBody, err := n.request(http.MethodPost, "https://music.163.com/weapi/v3/playlist/detail", strings.NewReader(params.Encode()))
	if err != nil {
		panic(err)
	}
	var payload = PlayListPayload{}
	json.Unmarshal(responseBody, &payload)
	return &payload
}

// 获取歌曲详情
func (n *Netease) GetSongDetail(ids []string, withSongUrl bool) *SongsDetail {
	// 组合参数，参数为json字符串
	idsByte, _ := json.Marshal(ids)
	var c []struct {
		Id string `json:"id"`
	}
	for _, id := range ids {
		c = append(c, struct {
			Id string `json:"id"`
		}{Id: id})
	}
	cByte, _ := json.Marshal(c)
	// 接口实际需要的参数
	paramsStruct := struct {
		Ids string `json:"ids"`
		C   string `json:"c"`
	}{
		Ids: string(idsByte),
		C:   string(cByte),
	}

	paramRaw, err := json.Marshal(paramsStruct)
	if err != nil {
		panic("json encode error")
	}
	crypto := ParamsCrypto{Plaint: string(paramRaw)}
	crypto.Encrypt()
	params := url.Values{}
	params.Set("params", crypto.Cipher.Param)
	params.Set("encSecKey", crypto.Cipher.EncSecKey)
	responseBody, err := n.request(http.MethodPost, "https://music.163.com/weapi/v3/song/detail", strings.NewReader(params.Encode()))
	if err != nil {
		panic(err)
	}
	songsDetail := SongsDetail{}
	json.Unmarshal(responseBody, &songsDetail)
	return &songsDetail
}

//// todo
//func (n *Netease) GetSongsUrl(ids []int) {
//	idsByte, _ := json.Marshal(ids)
//	paramsStruct := struct {
//		Ids string `json:"ids"`
//		Br  int    `json:"br"`
//	}{
//		Ids: string(idsByte),
//		Br:  999000,
//	}
//
//	pByte, err := json.Marshal(paramsStruct)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(pByte))
//	crypto := ParamsCrypto{Plaint: string(pByte)}
//	crypto.Encrypt()
//	params := url.Values{}
//	params.Set("params", crypto.Cipher.Param)
//	params.Set("encSecKey", crypto.Cipher.EncSecKey)
//	responseBody, err := n.request(http.MethodPost, "https://music.163.com/api/song/enhance/player/url", strings.NewReader(params.Encode()))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(responseBody))
//}

// 歌曲详情返回结构
type SongsDetail struct {
	Songs      []Song      `json:"songs"`
	Privileges []Privilege `json:"privileges"`
	Code       int         `json:"code"`
}

// 歌曲结构
type Song struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
	Pst  int    `json:"pst"`
	T    int    `json:"t"`
	Ar   []struct {
		ID    int           `json:"id"`
		Name  string        `json:"name"`
		Tns   []interface{} `json:"tns"`
		Alias []interface{} `json:"alias"`
	} `json:"ar"`
	Alia []interface{} `json:"alia"`
	Pop  float64       `json:"pop"`
	St   int           `json:"st"`
	Rt   string        `json:"rt"`
	Fee  int           `json:"fee"`
	V    int           `json:"v"`
	Crbt interface{}   `json:"crbt"`
	Cf   string        `json:"cf"`
	Dt   int           `json:"dt"`
	H    struct {
		Br   int     `json:"br"`
		Fid  int     `json:"fid"`
		Size int     `json:"size"`
		Vd   float64 `json:"vd"`
	} `json:"h"`
	M struct {
		Br   int     `json:"br"`
		Fid  int     `json:"fid"`
		Size int     `json:"size"`
		Vd   float64 `json:"vd"`
	} `json:"m"`
	L struct {
		Br   int     `json:"br"`
		Fid  int     `json:"fid"`
		Size int     `json:"size"`
		Vd   float64 `json:"vd"`
	} `json:"l"`
	A               interface{}   `json:"a"`
	Cd              string        `json:"cd"`
	No              int           `json:"no"`
	RtURL           interface{}   `json:"rtUrl"`
	Ftype           int           `json:"ftype"`
	RtUrls          []interface{} `json:"rtUrls"`
	DjID            int           `json:"djId"`
	Copyright       int           `json:"copyright"`
	SID             int           `json:"s_id"`
	Mark            int           `json:"mark"`
	OriginCoverType int           `json:"originCoverType"`
	NoCopyrightRcmd interface{}   `json:"noCopyrightRcmd"`
	Rtype           int           `json:"rtype"`
	Rurl            interface{}   `json:"rurl"`
	Mst             int           `json:"mst"`
	Cp              int           `json:"cp"`
	Mv              int           `json:"mv"`
	PublishTime     int64         `json:"publishTime"`
	Al              struct {
		ID     int           `json:"id"`
		Name   string        `json:"name"`
		PicURL string        `json:"picUrl"`
		Tns    []interface{} `json:"tns"`
		Pic    int64         `json:"pic"`
	} `json:"al,omitempty"`
}

// 特权结构
type Privilege struct {
	ID            int  `json:"id"`
	Fee           int  `json:"fee"`
	Payed         int  `json:"payed"`
	St            int  `json:"st"`
	Pl            int  `json:"pl"`
	Dl            int  `json:"dl"`
	Sp            int  `json:"sp"`
	Cp            int  `json:"cp"`
	Subp          int  `json:"subp"`
	Cs            bool `json:"cs"`
	Maxbr         int  `json:"maxbr"`
	Fl            int  `json:"fl"`
	Toast         bool `json:"toast"`
	Flag          int  `json:"flag"`
	PreSell       bool `json:"preSell"`
	PlayMaxbr     int  `json:"playMaxbr"`
	DownloadMaxbr int  `json:"downloadMaxbr"`
}

// 歌单详情返回结构
type PlayListPayload struct {
	Code          int         `json:"code"`
	RelatedVideos interface{} `json:"relatedVideos"`
	Playlist      struct {
		Subscribers []interface{} `json:"subscribers"`
		Subscribed  bool          `json:"subscribed"`
		Creator     struct {
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
		} `json:"creator"`
		Tracks   []Song `json:"tracks"`
		TrackIds []struct {
			ID  int         `json:"id"`
			V   int         `json:"v"`
			Alg interface{} `json:"alg"`
		} `json:"trackIds"`
		UpdateFrequency       interface{}   `json:"updateFrequency"`
		BackgroundCoverID     int           `json:"backgroundCoverId"`
		BackgroundCoverURL    interface{}   `json:"backgroundCoverUrl"`
		TitleImage            int           `json:"titleImage"`
		TitleImageURL         interface{}   `json:"titleImageUrl"`
		EnglishTitle          interface{}   `json:"englishTitle"`
		OpRecommend           bool          `json:"opRecommend"`
		CoverImgURL           string        `json:"coverImgUrl"`
		SpecialType           int           `json:"specialType"`
		Tags                  []interface{} `json:"tags"`
		SubscribedCount       int           `json:"subscribedCount"`
		CloudTrackCount       int           `json:"cloudTrackCount"`
		AdType                int           `json:"adType"`
		Privacy               int           `json:"privacy"`
		TrackUpdateTime       int64         `json:"trackUpdateTime"`
		NewImported           bool          `json:"newImported"`
		CoverImgID            int64         `json:"coverImgId"`
		UpdateTime            int64         `json:"updateTime"`
		CommentThreadID       string        `json:"commentThreadId"`
		TrackCount            int           `json:"trackCount"`
		TrackNumberUpdateTime int64         `json:"trackNumberUpdateTime"`
		PlayCount             int           `json:"playCount"`
		Description           interface{}   `json:"description"`
		Ordered               bool          `json:"ordered"`
		Status                int           `json:"status"`
		UserID                int           `json:"userId"`
		CreateTime            int64         `json:"createTime"`
		HighQuality           bool          `json:"highQuality"`
		Name                  string        `json:"name"`
		ID                    int           `json:"id"`
		ShareCount            int           `json:"shareCount"`
		CoverImgIDStr         string        `json:"coverImgId_str"`
		CommentCount          int           `json:"commentCount"`
	} `json:"playlist"`
	Urls       interface{} `json:"urls"`
	Privileges []Privilege `json:"privileges"`
}
