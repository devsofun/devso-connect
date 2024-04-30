package utils

import "testing"

func TestLuoguPasteChecker(t *testing.T) {
	result, _ := CheckParameterInURL(
		"https://www.luogu.com/paste/64cojga5", // 洛谷国际, $Host 更新
		"NqnpvkdSIBUfNqHVpuZQ",
	)
	expected := true
	if result != expected {
		t.Errorf("CheckParameterInURL() Result = %t; Want %t", result, expected)
	}
}
