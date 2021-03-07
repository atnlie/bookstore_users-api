package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(strPass string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(strPass))
	return hex.EncodeToString(hash.Sum(nil))
}
