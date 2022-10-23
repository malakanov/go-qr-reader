package libs

import (
	"bytes"
	"fmt"
	"image"
	"strings"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/multi/qrcode"
)

func Scan(b []byte) (string, string) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		msg := fmt.Sprintf("Failed to read image: %v", err)
		return "", msg
	}

	source := gozxing.NewLuminanceSourceFromImage(img)
	bin := gozxing.NewHybridBinarizer(source)
	bbm, err := gozxing.NewBinaryBitmap(bin)

	if err != nil {
		msg := fmt.Sprintf("Error during processing: %v", err)
		return "", msg
	}

	qrReader := qrcode.NewQRCodeMultiReader()
	result, err := qrReader.DecodeMultiple(bbm, nil)
	if err != nil {
		msg := fmt.Sprintf("Unable to decode QRCode: %v", err)
		return "", msg
	}
	strRes := []string{}
	for _, element := range result {
		strRes = append(strRes, element.String())
	}

	res := strings.Join(strRes, "\n")
	return res, ""
}
