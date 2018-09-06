package u

import "regexp"

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
