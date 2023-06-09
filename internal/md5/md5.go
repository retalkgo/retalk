package md5

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
