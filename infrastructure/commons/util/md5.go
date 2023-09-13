package util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	d := []byte(str)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func MD5WithSalt(str, salt string) (md5Str string, err error) {
	m5 := md5.New()
	_, err = m5.Write([]byte(str))
	if err != nil {
		return
	}
	_, err = m5.Write([]byte(salt))
	if err != nil {
		return
	}
	return hex.EncodeToString(m5.Sum(nil)), nil
}
