package gothumbor

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
)

type ThumborOptions struct {
	Width  int
	Height int
}

func GetThumborPath(key, imageUrl string, options ThumborOptions) (url string, err error) {
	var partial string
	if partial, err = GetUrlParts(imageUrl, options); err != nil {
		return
	}

	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(partial))
	message := hash.Sum(nil)
	url = base64.URLEncoding.EncodeToString(message)
	url = strings.Join([]string{url, partial}, "/")
	return url, err
}

func GetUrlParts(imageUrl string, options ThumborOptions) (urlPartial string, err error) {
	height := 0
	width := 0

	if options.Height != height {
		height = options.Height
	}

	if options.Width != width {
		width = options.Width
	}

	wXh := fmt.Sprintf("%dx%d", width, height)

	urlPartial = strings.Join([]string{wXh, imageUrl}, "/")
	return urlPartial, err
}
