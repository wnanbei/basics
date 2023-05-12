package user

import (
	"crypto/md5"
	"encoding/hex"
)

// ComparePassword 对比密码是否匹配
func ComparePassword(password, hash string) bool {
	return MD5(password, "ippool", 2) == hash
}

// MD5 多次加密，并加盐
func MD5(password string, salt string, iteration int) string {
	p := []byte(password)
	s := []byte(salt)

	h := md5.New()
	h.Write(s) // 先传入盐值
	h.Write(p)

	var res []byte
	res = h.Sum(nil)
	for i := 0; i < iteration-1; i++ {
		h.Reset()
		h.Write(res)
		res = h.Sum(nil)
	}
	return hex.EncodeToString(res)
}
