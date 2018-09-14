package u

import (
	"crypto/rand"
	"fmt"
	"os"
	"reflect"
	"regexp"
)

// IsValidIDCard check IDCard is valid or not.
func IsValidIDCard(IDCard string) bool {
	length := len(IDCard)

	if length != 15 && length != 18 {
		return false
	}

	if length == 15 {
		//验证15位身份证，15位的是全部数字
		if m, _ := regexp.MatchString(`^(\d{15})$`, IDCard); !m {
			return false
		}
		return true
	}

	if length == 18 {
		//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, IDCard); !m {
			return false
		}
		return true
	}

	return false
}

// IsDir judge whether dir is directory.
func IsDir(dir string) bool {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return fileInfo.Mode().IsDir()
}

// IsFile judge whether filename is file.
func IsFile(filename string) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return fileInfo.Mode().IsRegular()
}

// ToInt64 convert any numeric value to int64
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}

// GetRandomString generate random string
func GetRandomString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
