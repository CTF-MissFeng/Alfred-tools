package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/tjfoc/gmsm/sm3"

	"tools/util/consts"
)

func MD5(str1 string){
	srcCode := md5.Sum([]byte(str1))
	code := fmt.Sprintf("%x", srcCode)
	consts.Workflow.NewItem(code).
		Subtitle("md5").
		Icon(consts.HashIcon).
		Arg(code).
		Valid(true)
}

func SHA1(str1 string){
	h := sha1.New()
	h.Write([]byte(str1))
	bs := h.Sum(nil)
	hashCode := hex.EncodeToString(bs)
	consts.Workflow.NewItem(hashCode).
		Subtitle("sha1").
		Icon(consts.HashIcon).
		Arg(hashCode).
		Valid(true)
}

func SHA256(str1 string){
	h := sha256.New()
	h.Write([]byte(str1))
	bs := h.Sum(nil)
	hashCode := hex.EncodeToString(bs)
	consts.Workflow.NewItem(hashCode).
		Subtitle("sha256").
		Icon(consts.HashIcon).
		Arg(hashCode).
		Valid(true)
}

func SHA512(str1 string){
	h := sha512.New()
	h.Write([]byte(str1))
	bs := h.Sum(nil)
	hashCode := hex.EncodeToString(bs)
	consts.Workflow.NewItem(hashCode).
		Subtitle("sha512").
		Icon(consts.HashIcon).
		Arg(hashCode).
		Valid(true)
}

func SM3(str1 string){
	h := sm3.New()
	h.Write([]byte(str1))
	bs := h.Sum(nil)
	hashCode := hex.EncodeToString(bs)
	consts.Workflow.NewItem(hashCode).
		Subtitle("sm3").
		Icon(consts.HashIcon).
		Arg(hashCode).
		Valid(true)
}

func HashMain(str1 string){
	MD5(str1)
	SHA1(str1)
	SHA256(str1)
	SHA512(str1)
	SM3(str1)
}