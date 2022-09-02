package decode

import (
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"

	"tools/util/consts"
)

func Base64Decode(data string){
	decodeBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}
	decodeStr := string(decodeBytes)
	consts.Workflow.NewItem(decodeStr).
		Subtitle("base64").
		Icon(consts.EncodeIcon).
		Arg(decodeStr).
		Valid(true)
}

func UrlDecode(data string){
	enEscapeUrl, err := url.QueryUnescape(data)
	if err != nil{
		return
	}
	consts.Workflow.NewItem(enEscapeUrl).
		Subtitle("url").
		Icon(consts.EncodeIcon).
		Arg(enEscapeUrl).
		Valid(true)
}

func HexDecode(data string){
	hexData, err := hex.DecodeString(data)
	if err != nil{
		return
	}
	decode := string(hexData)
	consts.Workflow.NewItem(decode).
		Subtitle("hex").
		Icon(consts.EncodeIcon).
		Arg(decode).
		Valid(true)
}

func AsciiDecode(data string){
	strList := make([]string, 0)
	dataList := strings.Split(data, " ")
	for _, v := range dataList{
		int1, err := strconv.Atoi(v)
		if err != nil{
			return
		}
		tmp1 := string(rune(int1))
		strList = append(strList, tmp1)
	}
	decodeStr := strings.Join(strList, "")
	consts.Workflow.NewItem(decodeStr).
		Subtitle("ascii").
		Icon(consts.EncodeIcon).
		Arg(decodeStr).
		Valid(true)
}

func UnicodeDecode(data string){
	decodeStr, err := strconv.Unquote(strings.Replace(strconv.Quote(data), `\\u`, `\u`, -1))
	if err != nil{
		return
	}
	consts.Workflow.NewItem(decodeStr).
		Subtitle("unicode").
		Icon(consts.EncodeIcon).
		Arg(decodeStr).
		Valid(true)
}

func Main(data string){
	Base64Decode(data)
	UrlDecode(data)
	HexDecode(data)
	AsciiDecode(data)
	UnicodeDecode(data)
}
