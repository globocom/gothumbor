package gothumbor

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

type ThumborOptions struct {
	Width  int
	Height int
}

var (
	ErrorHeight = errors.New("Negative value height")
	ErrorWidth  = errors.New("Negative value width")
)

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
	if options.Height < 0 {
		return "", ErrorHeight
	}

	if options.Width < 0 {
		return "", ErrorWidth
	}

	wXh := fmt.Sprintf("%dx%d", options.Width, options.Height)
	urlPartial = strings.Join([]string{wXh, imageUrl}, "/")

	return urlPartial, nil
}
