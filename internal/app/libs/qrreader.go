package libs

import (
	"bytes"
	"fmt"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/multi/qrcode"
	"image"
	"strings"
)

func Scan(b []byte) (string, string) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		msg := fmt.Sprintf("failed to read image: %v", err)
		return "", msg
	}

	source := gozxing.NewLuminanceSourceFromImage(img)
	bin := gozxing.NewHybridBinarizer(source)
	bbm, err := gozxing.NewBinaryBitmap(bin)

	if err != nil {
		msg := fmt.Sprintf("error during processing: %v", err)
		return "", msg
	}

	qrReader := qrcode.NewQRCodeMultiReader()
	result, err := qrReader.DecodeMultiple(bbm, nil)
	if err != nil {
		msg := fmt.Sprintf("unable to decode QRCode: %v", err)
		return "", msg
	}
	strRes := []string{}
	for _, element := range result {
		strRes = append(strRes, element.String())
	}

	res := strings.Join(strRes, "\n")
	return res, ""
}
