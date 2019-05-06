package chardet

import (
	"io/ioutil"
	"strings"
)

// ToUTF8 - transform string with sourceCharset to utf-8
func ToUTF8(sourceCharset []string, source string) (string, error) {
	// 文件字符集校验
	charsetName := Mostlikein([]byte(source), sourceCharset)
	if charsetName != "utf-8" {
		charsetReader, err := NewReader(strings.NewReader(source), charsetName, []byte(source))
		if err != nil {
			return "", err
		}
		b, err := ioutil.ReadAll(charsetReader)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	return source, nil
}
