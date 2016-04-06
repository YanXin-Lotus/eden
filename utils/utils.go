package utils

import (
    "crypto/md5"
    "eden/config"
    "encoding/hex"
)

func MD5(raw string) (encrypt string) {
    hasher := md5.New()
    hasher.Write([]byte(raw))
    return hex.EncodeToString(hasher.Sum([]byte(config.Config.PWSecret)))
}