package api

import "testing"

var netease = Netease{}

func TestEncrypt(t *testing.T) {
	netease.GetPlayListDetail("38196761")
}

func TestNetease_GetSongDetail(t *testing.T) {
	ids := []string{"234252", "1422992414", "1446233390"}
	netease.GetSongDetail(ids, false)
}

func TestNetease_GetSongsUrl(t *testing.T) {
	ids := []int{234252, 1422992414, 1446233390}
	netease.GetSongsUrl(ids)
}
