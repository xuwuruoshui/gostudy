package split

import (
	"strings"
)

func Split(str, lable string) []string {

	// strings.Count分割后的长度
	result := make([]string, 0,strings.Count(str, lable)+1)

	i := strings.Index(str, lable)

	for i > -1 {
		result = append(result, str[:i])
		str = str[i+len(lable):]
		i = strings.Index(str, lable)
	}

	result = append(result, str)
	return result
}


