package utils

import (
	"crypto/md5"
	"eden/config"
	"encoding/hex"
	"errors"
	"regexp"
)

func MD5(raw string) (encrypt string) {
	hasher := md5.New()
	hasher.Write([]byte(raw))
	return hex.EncodeToString(hasher.Sum([]byte(config.Config.PWSecret)))
}

func CheckUsernamePermitted(name string) error {
	//username allow only letters and numbers
	m, err := regexp.Match("^[A-Za-z0-9]{4,20}$", []byte(name))
	if !m || err != nil {
		return errors.New("包含非法字符或者长度不合法(允许4-20位长度)")
	}

	return nil
}

func CheckPasswordPermitted(pw string) error {
	//password allow special symbol
	m, err := regexp.Match(`^[A-Za-z0-9!"#$%&'()*+,\-./:;<=>?@[\\\]^_{|}~]{6,32}$`, []byte(pw))
	if !m || err != nil {
		return errors.New("包含非法字符或者长度不合法(允许6-32位长度)")
	}

	return nil
}

//检查用户输入是否包含js脚本
func CheckContentPermitted(cont string) error {
	m, err := regexp.Match("<script>", []byte(cont))
	if !m || err != nil {
		return errors.New("包含非法字符/脚本")
	}

	return nil
}
