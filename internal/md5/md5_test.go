package md5

import "testing"

func TestMD5(t *testing.T) {
	data := MD5("Retalk")
	if data != "3402c46cde565fd412acaa4ac1bbcc9b" {
		t.Errorf("MD5函数校验失败")
	}
}