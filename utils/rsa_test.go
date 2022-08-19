package utils

import (
	"testing"
)

func TestRsaPriDecode(t *testing.T) {
	want := "123456"
	priEnContest := RsaPriEncode(want)
	t.Log("密文", priEnContest)
	got := RsaPriDecode(priEnContest)
	t.Log("want", priEnContest)
	if want != got {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
	t.Log("原文", got)
}
