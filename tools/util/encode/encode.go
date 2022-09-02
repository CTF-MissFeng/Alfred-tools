package encode

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"unicode"

	"tools/util/consts"
)

func Base64Encode(data string){
	encodeString := base64.StdEncoding.EncodeToString([]byte(data))
	consts.Workflow.NewItem(encodeString).
		Subtitle("base64").
		Icon(consts.EncodeIcon).
		Arg(encodeString).
		Valid(true)
}

func UrlEncode(data string){
	escapeUrl := url.QueryEscape(data)
	consts.Workflow.NewItem(escapeUrl).
		Subtitle("url").
		Icon(consts.EncodeIcon).
		Arg(escapeUrl).
		Valid(true)
}

func HexEncode(data string){
	hexStringData := hex.EncodeToString([]byte(data))
	consts.Workflow.NewItem(hexStringData).
		Subtitle("hex").
		Icon(consts.EncodeIcon).
		Arg(hexStringData).
		Valid(true)
}

func AsciiEncode(data string){
	asciiList := make([]string,0)
	for _, v := range []rune(data){
		intStr := fmt.Sprintf("%d", int(v))
		asciiList = append(asciiList, intStr)
	}
	encodeStr := strings.Join(asciiList, " ")
	consts.Workflow.NewItem(encodeStr).
		Subtitle("ascii").
		Icon(consts.EncodeIcon).
		Arg(encodeStr).
		Valid(true)
}

func UnicodeEncode(data string){
	DD := []rune(data)
	finallStr := ""
	for i := 0; i < len(DD); i++ {
		if unicode.Is(unicode.Scripts["Han"], DD[i]) {
			textQuoted := strconv.QuoteToASCII(string(DD[i]))
			finallStr += textQuoted[1 : len(textQuoted)-1]
		} else {
			h := fmt.Sprintf("%x",DD[i])
			finallStr += "\\u" + isFullFour(h)
		}
	}
	consts.Workflow.NewItem(finallStr).
		Subtitle("unicode").
		Icon(consts.EncodeIcon).
		Arg(finallStr).
		Valid(true)
}

func isFullFour(str string) (string) {
	if len(str) == 1 {
		str = "000" + str
	} else if len(str) == 2 {
		str = "00" + str
	} else if len(str) == 3 {
		str = "0" + str
	}
	return str
}

func Main(data string){
	Base64Encode(data)
	UrlEncode(data)
	HexEncode(data)
	AsciiEncode(data)
	UnicodeEncode(data)
}