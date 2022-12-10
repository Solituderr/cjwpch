package service

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

const (
	VAL   = 0x3FFFFFFFF
	INDEX = 0x0000003E
)

var (
	alphabet = []byte("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func Transform(originURL string) (string, error) {
	md5Str := getMd5Str(originURL)
	//var hexVal int64
	var tempVal int64
	var result [4]string
	var tempUri []byte
	var output string
	for i := 0; i < 4; i++ {
		tempSubStr := md5Str[i*8 : (i+1)*8]
		hexVal, err := strconv.ParseInt(tempSubStr, 16, 64)
		if err != nil {
			return "fail", err
		}
		tempVal = int64(VAL) & hexVal
		var index int64
		tempUri = []byte{}
		for i := 0; i < 6; i++ {
			index = INDEX & tempVal
			tempUri = append(tempUri, alphabet[index])
			tempVal = tempVal >> 5
		}
		result[i] = string(tempUri)
		output += result[i]
	}
	output = "https://shorturl.com/" + output
	return output, nil
}

// generate md5 checksum of URL in hex format
func getMd5Str(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	c := m.Sum(nil)
	return hex.EncodeToString(c)
}

/*
func main() {
	var url string
	url = "https://translate.google.com/?sl=auto&tl=zh-CN&text=Reverse%20Engineered%20ChatGPT%20by%20OpenAI.%20Extensible%20for%20chatbots%20etc.&op=translate"
	url, _ = Transform(url)
	fmt.Println(url)
}
*/
