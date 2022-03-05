package cmb

import (
	"github.com/axgle/mahonia"
	"strings"
)


var gbkDecoder = mahonia.NewDecoder("gbk")
var utf8Decoder = mahonia.NewDecoder("utf-8")

func convertGBK(source string) string {
	resStr:= gbkDecoder.ConvertString(source)
	_, resBytes, _ := utf8Decoder.Translate([]byte(resStr), true)
	return string(resBytes)
}

func convertAccount(source string) string {
	return strings.Split(strings.Split(strings.Split(source, "[")[1], " ")[0], ":")[1]
}